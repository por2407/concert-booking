package services

import (
	"errors"
	"time"

	"github.com/ticket-backend/internal/adapters/repositories"
	"github.com/ticket-backend/internal/core/domain"
	"gorm.io/gorm"
)

type BookingService struct {
	db          *gorm.DB // เราต้องใช้ DB เพื่อสั่ง Start Transaction
	seatRepo    *repositories.PostgresSeatRepo
	bookingRepo *repositories.PostgresBookingRepo
}

// รับ Repo เข้ามา (สังเกตว่าผมใช้ Type จริง (struct) แทน Interface ชั่วคราวเพื่อให้ง่ายต่อการส่ง Transaction)
func NewBookingService(db *gorm.DB, seatRepo *repositories.PostgresSeatRepo, bookingRepo *repositories.PostgresBookingRepo) *BookingService {
	return &BookingService{
		db:          db,
		seatRepo:    seatRepo,
		bookingRepo: bookingRepo,
	}
}

func (s *BookingService) CreateBooking(userID uint, seatIDs []uint) (*domain.Booking, error) {
	var newBooking domain.Booking

	// 1. เริ่ม Transaction (Database Transaction)
	// กฎ: ถ้า Error ตรงไหนก็ตาม ข้อมูลทุกอย่างจะย้อนกลับเหมือนไม่เคยเกิดขึ้น (Rollback)
	err := s.db.Transaction(func(tx *gorm.DB) error {

		totalAmount := 0.0
		var bookingItems []domain.BookingItem

		// 2. วนลูปตรวจสอบและล็อคที่นั่งทีละตัว
		for _, seatID := range seatIDs {
			// 🔥 จุดตาย: เรียกใช้ GetSeatWithLock (SELECT ... FOR UPDATE)
			// ส่ง tx เข้าไป เพื่อบอกว่า "ให้ทำใน Transaction นี้นะ"
			seat, err := s.seatRepo.GetSeatWithLock(tx, seatID)
			if err != nil {
				return err // หาไม่เจอ
			}

			// 3. เช็คว่าว่างไหม? (Double Booking Check)
			if seat.Status != "AVAILABLE" {
				return errors.New("seat " + seat.RowLabel + " is not available")
			}

			// 4. อัปเดตสถานะที่นั่ง (Lock ไว้ก่อน)
			seat.Status = "LOCKED"
			seat.LockedBy = &userID
			expireTime := time.Now().Add(10 * time.Minute) // จองไว้ 10 นาที
			seat.LockExpiresAt = &expireTime

			if err := tx.Save(&seat).Error; err != nil {
				return err
			}

			// คำนวณเงิน
			totalAmount += seat.Price

			// เตรียมข้อมูล Item
			bookingItems = append(bookingItems, domain.BookingItem{
				SeatID: seat.ID,
			})
		}

		// 5. สร้างใบจอง (Booking Header)
		newBooking = domain.Booking{
			UserID:      userID,
			EventID:     1, // สมมติว่างาน ID 1 ไปก่อน (จริงๆ ต้องเช็คว่า seat มาจากงานไหน)
			TotalAmount: totalAmount,
			Status:      domain.BookingPending,
			Items:       bookingItems, // GORM จะสร้าง BookingItem ให้อัตโนมัติเพราะเราใส่ Relation ไว้
		}

		if err := s.bookingRepo.Create(tx, &newBooking); err != nil {
			return err
		}

		return nil // ถ้า return nil -> Transaction จะ Commit (บันทึกจริง)
	})

	if err != nil {
		return nil, err
	}

	return &newBooking, nil
}
