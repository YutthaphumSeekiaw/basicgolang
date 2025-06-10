package core

// Secondary port

// Secondary Port (ส่วนที่จะต้องส่งต่อไปยังส่วนของ database) = เก็บไว้ที่ order_repository.go โดย
// implement OrderRepository สำหรับ spec ข้อมูลสำหรับการรับข้อมูลเพื่อไป save ลง database (ซึ่งก็คือข้อมูล order)

type OrderRepository interface {
	Save(order Order) error
}
