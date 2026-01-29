package services

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9" // import redis
	"github.com/ticket-backend/internal/adapters/repositories"
	"github.com/ticket-backend/internal/core/domain"
	"gorm.io/gorm"
)

type bookingService struct {
	db          *gorm.DB
	redis       *redis.Client // เพิ่ม Redis เข้ามา
	seatRepo    *repositories.PostgresSeatRepo
	bookingRepo *repositories.PostgresBookingRepo
}

// รับ Redis Client เข้ามาด้วย
func NewBookingService(db *gorm.DB, rdb *redis.Client, seatRepo *repositories.PostgresSeatRepo, bookingRepo *repositories.PostgresBookingRepo) *bookingService {
	return &bookingService{
		db:          db,
		redis:       rdb,
		seatRepo:    seatRepo,
		bookingRepo: bookingRepo,
	}
}

func (s *bookingService) CreateBooking(userID uint, seatIDs []uint) (*domain.Booking, error) {
	ctx := context.Background()

	// --- 🛡️ ด่านที่ 1: Redis Guard (คัดกรองคน 99.9% ออกไปตรงนี้) ---
	// วนลูปเช็คว่ามีใครจองที่นั่งพวกนี้ใน Redis อยู่ไหม
	for _, seatID := range seatIDs {
		lockKey := fmt.Sprintf("lock:seat:%d", seatID)

		// SETNX: ถ้ายังไม่มี key นี้ ให้สร้างและ return true (ชนะ)
		// ถ้ามี key นี้อยู่แล้ว return false (แพ้)
		// ตั้งเวลาหมดอายุ 10 วินาที (เผื่อระบบล่ม key จะได้ไม่ค้างตลอดไป)
		success, err := s.redis.SetNX(ctx, lockKey, userID, 10*time.Second).Result()

		if err != nil {
			return nil, err
		}
		if !success {
			// ถ้าแพ้ใน Redis ให้ดีดออกเลย ไม่ต้องไปกวน DB
			return nil, fmt.Errorf("seat %d is currently being booked by someone else (Redis Block)", seatID)
		}

		// *ข้อควรระวัง: ในระบบจริง ถ้าจองหลายที่แล้วติดใบหลังๆ ต้องไปไล่ลบ Redis Key ของใบแรกๆ ออกด้วย (Compensating Transaction)
		// แต่นี่เอา Concept หลักก่อนครับ
	}
	// -----------------------------------------------------------

	// --- 🏰 ด่านที่ 2: Database Transaction (คนชนะเท่านั้นที่มาถึงตรงนี้) ---
	// โค้ดส่วนนี้เหมือนเดิม 100% เพราะเราต้อง Lock DB เพื่อความชัวร์ (Consistency)
	// แต่ภาระ DB จะเหลือน้อยมากๆ เพราะ Redis กั้นคนส่วนใหญ่ไว้แล้ว

	var newBooking domain.Booking
	err := s.db.Transaction(func(tx *gorm.DB) error {
		// ... (โค้ดเดิมทั้งหมดของคุณ อยู่ในนี้) ...
		// ...
		// ...
		return nil
	})

	// --- 🧹 ด่านที่ 3: Cleanup ---
	// ไม่ว่าจะจองสำเร็จหรือไม่ ควรลบ Key ใน Redis ออก (หรือปล่อยให้มัน Expire เองก็ได้ถ้าขี้เกียจ แต่ลบเลยดีกว่า)
	for _, seatID := range seatIDs {
		s.redis.Del(ctx, fmt.Sprintf("lock:seat:%d", seatID))
	}

	if err != nil {
		return nil, err
	}
	return &newBooking, nil
}
