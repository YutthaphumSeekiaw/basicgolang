package adapter

import (
	core "gohexarch/cores"

	"gorm.io/gorm"
)

// Secondary adapter

// Secondary Adapter (adapter สำหรับการส่งต่อข้อมูลเข้า database) = เก็บไว้ที่ gorm_adapter.go
// ฝั่ง Adapter ทำการสร้าง struct GormOrderRepository เป็น Secondary adapter โดยทำการ implement
// ตาม OrderRepository (ที่เป็น interface ของ Secondary port) และทำการเก็บ logic ที่ใช้สำหรับการพูดคุยกับ database เอาไว้
// ซึ่งในที่นี้คนที่ทำหน้าที่ช่วยคุยใน Secondary adapter ก็คือ library GORM จะเป็นคนไปคุยกับ database
//
//	ให้โดยการแปลงข้อมูลที่ส่งมา (โดยเรียกใช้งานจาก logic ของ Business logic) ทำการส่งเข้า database โดยตรงไป
type GormOrderRepository struct {
	db *gorm.DB
}

func NewGormOrderRepository(db *gorm.DB) core.OrderRepository {
	return &GormOrderRepository{db: db}
}

func (r *GormOrderRepository) Save(order core.Order) error {
	if result := r.db.Create(&order); result.Error != nil {
		// Handle database errors
		return result.Error
	}
	return nil
}

// จาก code

// มี struct GormOrderRepository ที่จะทำการเก็บ function ของการทำงาน database ไว้ (ซึ่งก็คือ method Save())
// NewGormOrderRepository function สำหรับการ instance ตัว GormOrderRepository โดยทำการรับ db (ซึ่งก็คือ database ที่จะทำการเชื่อมต่อ) เข้ามาผ่าน struct ที่มีการกำหนด spec เอาไว้
