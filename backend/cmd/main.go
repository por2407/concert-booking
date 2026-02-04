package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ticket-backend/internal/adapters/handlers"
	"github.com/ticket-backend/internal/adapters/middleware"
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

	// 2. Setup Redis (เพิ่ม)
	rdb, err := infrastructure.NewRedisClient()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
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
	userRepo := repositories.NewPostgresUserRepo(db)
	eventRepo := repositories.NewPostgresEventRepo(db)
	seatRepo := repositories.NewPostgresSeatRepo(db)
	bookingRepo := repositories.NewPostgresBookingRepo(db)
	txManager := repositories.NewGormTransactionManager(db)

	// 4. เตรียมผู้จัดการ (Service)
	jwtSecret := "super-secret-key" // ในงานจริงควรใช้แปรสภาพแวดล้อม (Env Var)
	authService := services.NewAuthService(userRepo, jwtSecret)
	svc := services.NewEventService(txManager, eventRepo, seatRepo)
	bookingService := services.NewBookingService(txManager, rdb, seatRepo, bookingRepo)

	// Handlers (สร้างคนรับแขก)
	authHandler := handlers.NewAuthHandler(authService)
	eventHandler := handlers.NewEventHandler(svc)
	bookingHandler := handlers.NewBookingHandler(bookingService)

	// 6. รัน Server
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Ticket System Running...")
	})

	api := app.Group("/api")

	// --- Auth Routes ---
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	// --- Public Routes ---
	api.Get("/events", eventHandler.GetAllEvents)
	api.Get("/events/:id", eventHandler.GetEvent)

	// --- Admin Routes (Protect หรือไม่ตามสะดวก แต่อันนี้ตัวอย่าง) ---
	api.Post("/events/create", eventHandler.CreateEvent)

	// --- Protected Routes (ต้อง Login ก่อน) ---
	protected := api.Group("/", middleware.JWTMiddleware(jwtSecret))
	protected.Post("/bookings", bookingHandler.CreateBooking)
	protected.Post("/bookings/confirm", bookingHandler.ConfirmBooking)
	protected.Get("/bookings/history", bookingHandler.GetHistory)

	// Start Server
	fmt.Println("Server listening on :8080")
	app.Listen(":8080")
}
