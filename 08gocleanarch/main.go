package main

import (
	adapters "gocleanarch/adapter"
	"gocleanarch/entities"
	"gocleanarch/usecases"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := db.AutoMigrate(&entities.Order{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	orderRepo := adapters.NewGormOrderRepository(db)
	orderService := usecases.NewOrderService(orderRepo)
	orderHandler := adapters.NewHttpOrderHandler(orderService)

	app.Post("/order", orderHandler.CreateOrder)

	log.Fatal(app.Listen(":8000"))
}

// อธิบายตาม code คือ

// ขั้นแรกกำหนด config ของ GORM (ส่วน database เชื่อมไปยัง sqlite) และ Fiber (ทำ HTTP Server)
// ต่อมาทำการส่ง db (ตัวแทนของ GORM) เข้าไปยัง adapters.NewGormOrderRepository เพื่อทำการแปลงเป็น Order Repository สำหรับใช้งานใน Order Service
// ส่ง Order Repository เข้าไปยัง usecases.NewOrderService(orderRepo) เพื่อทำการแปลงเป็น Order Service ที่เป็นตัวแทนสำหรับคุยในขั้นของ Use case
// สุดท้ายส่ง Order Service เข้าไปใน adapters.NewHttpOrderHandler(orderService) เพื่อทำการแปลงเป็น Order HTTP Handler เป็นตัวแทนของการคุย Request นี้ออกมา โดย Handler นั้นก็จะทำการเรียกใช้จาก Service ที่มีการส่งเข้าไปสำหรับการทำ logic ภายใน Handler นั้น
// และที่เหลือก็ทำการผูก orderHandler.CreateOrder เข้ากับ endpoint POST /order เพื่อให้สามารถเรียกใช้งาน CreateOrder() จากใน Adapter ได้
