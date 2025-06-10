package adapter

import (
	"fmt"
	"gohexarch/core"

	"github.com/gofiber/fiber/v2"
)

// Primary adapter

// Primary Adapter (adapter สำหรับการ handle input ที่เข้ามา) = เก็บไว้ที่ http_adapter.go
// ฝั่ง Adapter ทำการสร้าง struct HttpOrderHandler ขึ้นมา โดยทำการ implement ตาม Primary port (ซึ่งก็คือ OrderService ที่สร้างไว้)
// โดย Adapter นี้มีหน้าที่ในการ แปลง HTTP Request ที่เข้ามา เพื่อทำการเตรียมส่งเข้า OrderService (เพื่อให้ OrderService ทำการส่งต่อไปยัง Business logic ต่อ)
// ซึ่งโจทย์ของ Adapter นั้นมีเพียงแค่ปั้น data ให้ตรงตาม spec ของ OrderService เท่านั้น
type HttpOrderHandler struct {
	service core.OrderService
}

func NewHttpOrderHandler(service core.OrderService) *HttpOrderHandler {
	return &HttpOrderHandler{service: service}
}

func (h *HttpOrderHandler) CreateOrder(c *fiber.Ctx) error {
	var order core.Order
	if err := c.BodyParser(&order); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := h.service.CreateOrder(order); err != nil {
		// Return an appropriate error message and status code
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}

// จาก code

// มี HttpOrderHandler สำหรับเป็น struct กำหนดรูปแบบของ instance ของ Primary Adapter ที่จะต้องมีการส่งเข้ามา
// NewHttpOrderHandler ทำการ instance โดยรับ service ที่จะทำการส่งต่อเข้ามา (ในทีนี้คือ OrderService)
// ส่วน CreateOrder ก็จะเป็น method ของ HttpOrderHandler เพื่อให้เรียกใช้สำหรับ endpoint ของการ CreateOrder (สร้างขึ้นมาเพื่อให้ client ที่มีการเรียกใช้เป็น Primary Adapter สามารถเรียกใช้งานได้)
