package adapters

import (
	"gocleanarch/entities"
	"gocleanarch/usecases"

	"github.com/gofiber/fiber/v2"
)

// Adapter
// ต่อมา Adapter คือส่วนของการแปลงข้อมูลเพื่อทำการส่งไปยัง Use case เพื่อให้ Use case สามารถจัดการต่อได้ถูกได้จะมีทั้งหมด 2 ส่วน (โดยจะแยกออกเป็น 2 files) คือ

// gorm_order_repository.go (GORM Order Repository) คือส่วนที่จะเก็บ function สำหรับจัดการ database เอาไว้ โดยจะเก็บคำสั่งที่เกี่ยวข้องกับการ query database ไว้ โดยจะต้อง implement ตาม spec ของ interface ที่กำหนดไว้ใน Order Service
// http_order_handler.go (HTTP Order Handler) คือส่วนที่จะเก็บ function สำหรับจัดการ data ที่ผ่านเข้ามาทาง HTTP Request โดยทำการแปลงข้อมูลให้ถูกต้องตาม Pattern ของ Entities เพื่อส่งให้ Order Service สามารถจัดการต่อได้

type HttpOrderHandler struct {
	orderUseCase usecases.OrderUseCase
}

func NewHttpOrderHandler(useCase usecases.OrderUseCase) *HttpOrderHandler {
	return &HttpOrderHandler{orderUseCase: useCase}
}

func (h *HttpOrderHandler) CreateOrder(c *fiber.Ctx) error {
	var order entities.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := h.orderUseCase.CreateOrder(order); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}

// อธิบาย code เพิ่มเติม

// gorm_order_repository.go ทำการ implement Save() โดยทำการใส่ query ที่เกี่ยวข้องกับการจัดการ Order เข้าไป (ซึ่งก็คือคำสั่ง Gorm)
// http_order_handler.go ทำการ implement HttpOrderHandler โดยเป็นการรับ ตัวแทนของ Use case มาเพื่อใช้คำสั่งภายใน
//  Order Service (คำสั่งสำหรับการสร้าง Order) และทำการสร้าง method CreateOrder() เพื่อใช้สำหรับการเรียกใช้งานจากฝั่งของ
// HTTP Request ออกมา โดยจะเรียกไปยัง Order Service และทำการแปลงข้อมูลจาก HTTP Request มาเป็นข้อมูลของ Order Entity
// เพื่อให้สามารถใช้งานที่ Order Service ได้
