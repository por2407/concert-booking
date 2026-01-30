package services

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/ticket-backend/internal/core/domain"
	"github.com/ticket-backend/internal/core/ports"
	"gorm.io/gorm"
)

type bookingService struct {
	db          *gorm.DB
	redis       *redis.Client
	seatRepo    ports.SeatRepository
	bookingRepo ports.BookingRepository
}

func NewBookingService(db *gorm.DB, rdb *redis.Client, seatRepo ports.SeatRepository, bookingRepo ports.BookingRepository) *bookingService {
	return &bookingService{
		db:          db,
		redis:       rdb,
		seatRepo:    seatRepo,
		bookingRepo: bookingRepo,
	}
}

// 1. เส้นทางแรก: จองรอชำระเงิน (LOCKED)
func (s *bookingService) CreatePendingBooking(userID uint, seatIDs []uint) (*domain.Booking, error) {
	ctx := context.Background()

	// --- 🛡️ ด่านที่ 1: Redis Guard ---
	for _, seatID := range seatIDs {
		lockKey := fmt.Sprintf("lock:seat:%d", seatID)
		success, err := s.redis.SetNX(ctx, lockKey, userID, 10*time.Second).Result()
		if err != nil {
			return nil, err
		}
		if !success {
			return nil, fmt.Errorf("seat %d is currently being booked by someone else (Redis Block)", seatID)
		}
	}

	var newBooking domain.Booking
	err := s.db.Transaction(func(tx *gorm.DB) error {
		var totalAmount float64
		var bookingItems []domain.BookingItem
		var firstEventID uint
		expireTime := time.Now().Add(10 * time.Minute) // ล็อกใน DB ไว้ 10 นาที

		for _, seatID := range seatIDs {
			// Lock ที่นั่งใน DB
			seat, err := s.seatRepo.GetSeatWithLock(tx, seatID)
			if err != nil {
				return fmt.Errorf("seat %d not found", seatID)
			}

			// เฉพาะที่นั่งว่างเท่านั้นที่จองได้
			if seat.Status != "AVAILABLE" {
				return fmt.Errorf("seat %d is not available (status: %s)", seatID, seat.Status)
			}

			firstEventID = seat.EventID
			totalAmount += seat.Price

			// อัปเดตสถานะเป็น LOCKED
			seat.Status = "LOCKED"
			seat.LockedBy = &userID
			seat.LockExpiresAt = &expireTime
			if err := tx.Save(seat).Error; err != nil {
				return err
			}

			bookingItems = append(bookingItems, domain.BookingItem{
				SeatID: seat.ID,
			})
		}

		newBooking = domain.Booking{
			UserID:      userID,
			EventID:     firstEventID,
			TotalAmount: totalAmount,
			Status:      "PENDING",
			Items:       bookingItems,
		}

		if err := s.bookingRepo.Create(tx, &newBooking); err != nil {
			return err
		}
		return nil
	})

	// Cleanup Redis Lock
	for _, seatID := range seatIDs {
		s.redis.Del(ctx, fmt.Sprintf("lock:seat:%d", seatID))
	}

	if err != nil {
		return nil, err
	}
	return &newBooking, nil
}

// 2. เส้นทางที่สอง: ยืนยันการชำระเงิน (SOLD)
func (s *bookingService) ConfirmPayment(bookingID uint) (*domain.Booking, error) {
	var booking *domain.Booking
	err := s.db.Transaction(func(tx *gorm.DB) error {
		// 1. ดึงใบจองมาเช็ค
		var err error
		booking, err = s.bookingRepo.GetByID(bookingID)
		if err != nil {
			return fmt.Errorf("booking %d not found", bookingID)
		}

		if booking.Status != "PENDING" {
			return fmt.Errorf("booking %d is already %s", bookingID, booking.Status)
		}

		// 2. อัปเดตสถานะที่นั่งทุกลูกค้าที่อยู่ในใบจองนี้เป็น SOLD
		for _, item := range booking.Items {
			seat, err := s.seatRepo.GetSeatWithLock(tx, item.SeatID)
			if err != nil {
				return err
			}

			seat.Status = "SOLD"
			seat.LockExpiresAt = nil // เคลียร์วันหมดอายุ
			if err := tx.Save(seat).Error; err != nil {
				return err
			}
		}

		// 3. อัปเดตสถานะใบจองเป็น PAID
		booking.Status = "PAID"
		if err := s.bookingRepo.Update(tx, booking); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return booking, nil
}
