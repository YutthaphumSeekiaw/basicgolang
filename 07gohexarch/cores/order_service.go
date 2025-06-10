package core

import "errors"

// Primary port

// Primary Port (ส่วนที่ต้องรับจากภายนอก) = เก็บไว้ที่ order_service.go โดย
// implement OrderService เป็น Port ระบุการเชื่อมต่อเอาไว้ว่า ถ้าจะส่งข้อมูลเข้ามาต้องส่งข้อมูล order หน้าตาแบบไหนมา
// โดยเราจะสร้าง order.go สำหรับการเก็บ schema ของ data เอาไว้ เพื่อใช้สำหรับกำหนด spec ทั้ง Port, Adapter, Business Logic
//
//	ว่า data ที่ใช้สื่อสารกันมีหน้าตาเป็นแบบไหน
type OrderService interface {
	CreateOrder(order Order) error
}

// โดย เราจะทำการสร้าง function สำหรับจัดการ Order ขึ้นมา โดยใช้ function จาก
// Repository มาจัดการผูก logic (เป็น business function ขึ้นมา) สำหรับจัดการตัว Order
type orderServiceImpl struct {
	repo OrderRepository
}

func NewOrderService(repo OrderRepository) OrderService {
	return &orderServiceImpl{repo: repo}
}

func (s *orderServiceImpl) CreateOrder(order Order) error {
	if order.Total <= 0 {
		return errors.New("total must be positive")
	}
	// Business logic...
	if err := s.repo.Save(order); err != nil {
		return err
	}
	return nil
}

// โดย

// ตัว struct orderServiceImpl จะทำการเก็บ logic ของการจัดการ Order เอาไว้ ชื่อ CreateOrder(order Order) โดยจะทำการคุยกับ OrderRepository (ที่จะส่งเข้ามาจาก Adapter อีกทีตาม spec ที่ Port กำหนด) เพื่อทำการไปคุยให้สร้าง Order จาก Database ออกมา
// ส่วน function NewOrderService ไว้สำหรับการ instance ตัวแปรตอนที่มีการสร้างตัวแปรตาม struct ของ orderServiceImpl (เป็นตัวที่ใช้รับ Repository ที่ส่งจาก adapter เข้ามา)
