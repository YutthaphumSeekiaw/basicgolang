package main

import (
	"gohexarch/adapter"
	"gohexarch/core"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ทำการกำหนด config database (ในทีนี้ขอใช้ sqlite เพื่อความรวดเร็วในการ implement) และ fiber (port สำหรับการ run http server)
// นำตัวแปร config ของ database ไล่ใส่ไปตั้งแต่ Secondary Adapter (ผ่าน Port) > Business logic (Service) > Primary Adapter (ผ่าน Port) โดยถ้าเทียบกับโจทย์ที่เรา implement ไปก็จะเป็น GormOrderRepository > OrderService > HttpOrderHandler

func main() {
	app := fiber.New()

	// Initialize the database connection
	db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&core.Order{})

	// Set up the core service and adapters
	orderRepo := adapter.NewGormOrderRepository(db)
	orderService := core.NewOrderService(orderRepo)
	orderHandler := adapter.NewHttpOrderHandler(orderService)

	// Define routes
	app.Post("/order", orderHandler.CreateOrder)

	// Start the server
	app.Listen(":8000")
}
