package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ticket-backend/internal/adapters/repositories"
	"github.com/ticket-backend/internal/core/domain"
	"github.com/ticket-backend/internal/core/services"
	"github.com/ticket-backend/internal/infrastructure"
)

func main() {
	// 1. ต่อ DB
	db, err := infrastructure.NewPostgresDB()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// --- [เพิ่มตรงนี้] สร้าง ENUM ก่อน Migrate ---
	// เช็คว่าถ้ายังไม่มี Type นี้ ให้สร้างใหม่ (กัน Error เวลา Restart)
	db.Exec(`DO $$ BEGIN
		CREATE TYPE seat_status AS ENUM ('AVAILABLE', 'LOCKED', 'SOLD');
	EXCEPTION
		WHEN duplicate_object THEN null;
	END $$;`)

	db.Exec(`DO $$ BEGIN
		CREATE TYPE booking_status AS ENUM ('PENDING', 'PAID', 'CANCELLED');
	EXCEPTION
		WHEN duplicate_object THEN null;
	END $$;`)
	// ----------------------------------------

	// 2. สร้างตาราง (Migration)
	db.AutoMigrate(
		&domain.User{},
		&domain.Event{},
		&domain.Seat{},
		&domain.Booking{},
		&domain.BookingItem{},
	)

	// 3. เตรียมคนงาน (Repo)
	eventRepo := repositories.NewPostgresEventRepo(db)
	seatRepo := repositories.NewPostgresSeatRepo(db)

	// 4. เตรียมผู้จัดการ (Service)
	svc := services.NewEventService(eventRepo, seatRepo)

	// 5. สั่งเสกข้อมูล (ทำงานเบื้องหลัง)
	go func() {
		fmt.Println("⏳ Seeding data...")
		svc.SeedData()
	}()

	// 6. รัน Server
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("System Ready!")
	})
	app.Listen(":8080")
}
