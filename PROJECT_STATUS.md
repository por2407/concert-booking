# Project Status Report - Concert Booking System

## 🚀 สถานะปัจจุบัน (Current Status)

โปรเจกต์อยู่ในขั้นตอนการพัฒนา Backend Core และการตั้งค่า Infrastructure (Docker, PostgreSQL, Redis) โดยขณะนี้ได้แก้ไขข้อผิดพลาดในส่วนของการเชื่อมต่อ Database และ Repository เรียบร้อยแล้ว

### สิ่งที่แก้ไขล่าสุด:

- [x] **Module Alignment**: เปลี่ยนชื่อโมดูลใน `go.mod` เป็น `github.com/ticket-backend` ให้ตรงกับ import ในโค้ด ช่วยให้ Air build ผ่าน
- [x] **Database Initialization**: แก้ไข Syntax Error ใน `main.go` และเพิ่ม Error Handling เพื่อไม่ให้ระบบค้างหากต่อ DB ไม่ได้
- [x] **Repository Implementation**: เพิ่มไฟล์ `postgres_seat_repo.go` ที่ขาดหายไปเพื่อให้ระบบสามารถจัดการข้อมูลที่นั่งได้
- [x] **Domain Alignment**: แก้ไขชื่อ Field ใน `event_service.go` ให้ตรงกับ `domain` (เช่น `Date` -> `DateTime`, `Row` -> `RowLabel`)
- [x] **SQL Enum Creation**: ระบบจะสร้าง custom enum `seat_status` และ `booking_status` ใน Postgres โดยอัตโนมัติก่อนเริ่ม Migration

---

## ❓ ทำไม GORM ไม่สร้าง Database?

**คำตอบ**: GORM มีหน้าที่จัดการ **Table/Index (Schema)** ภายใน Database เท่านั้น **ไม่สามารถสร้างตัว Database (เช่น `CREATE DATABASE ticket_db`) เองได้**

ในโปรเจกต์นี้ เราจัดการเรื่องนี้ผ่าน 2 ช่องทาง:

1. **Docker Compose**: ในไฟล์ `docker-compose.yml` มีการกำหนด `POSTGRES_DB: ticket_db` ซึ่ง Postgres Image จะสร้าง DB นี้ให้โดยอัตโนมัติเมื่อ Container เริ่มทำงานครั้งแรก
2. **Manual**: หากคุณรัน Backend นอก Docker คุณต้องเข้าไปสร้าง Database ชื่อ `ticket_db` ใน Postgres ด้วยตัวเองก่อน GORM ถึงจะเชื่อมต่อและสร้างตารางให้ได้

---

## 📂 โครงสร้างโปรเจกต์ (Project Structure)

โปรเจกต์นี้ใช้โครงสร้างแบบ **Hexagonal Architecture (Clean Architecture)** แบ่งส่วนชัดเจน:

```text
backend/
├── cmd/
│   └── main.go                 # จุดเริ่มต้นแอปพลิเคชัน (Entry Point)
├── internal/
│   ├── core/
│   │   ├── domain/             # (Entity) โมเดลข้อมูลหลัก (User, Event, Seat)
│   │   ├── ports/              # (Interface) ตัวกำหนดสัญญาการทำงานของ Repo/Service
│   │   └── services/           # (Use Case) Logic การทำงานหลักของระบบ
│   ├── adapters/
│   │   └── repositories/       # (Infrastructure) การคุยกับ Database จริงๆ (GORM/Postgres)
│   └── infrastructure/         # การตั้งค่าพื้นฐาน (Database connection)
├── Dockerfile.dev              # สำหรับรัน Backend ใน Docker
└── docker-compose.yml          # จัดการ Service ทั้งหมด (DB, Redis, PgAdmin, API)
```

---

## 🛠️ ขั้นตอนถัดไป

1. สร้าง API Endpoints เพิ่มเติม (Booking, Auth)
2. พัฒนาส่วน Frontend เพื่อเชื่อมต่อกับ API
3. ทดสอบการจองพร้อมกัน (Concurrency Test) ด้วย Redis
