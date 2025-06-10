package usecases

import "gocleanarch/entities"

// Usecase
// ส่วนต่อมา Use case คือส่วนของการใช้งาน ซึ่งในที่นี่คือ function createOrder สำหรับการสร้าง Order ซึ่งจะเก็บเอาไว้ใน Order Service

// โดย Order Service นั้นก็จะอ้างอิงโครงสร้างของ Order ตาม Order Entity ที่อยู่ใน layer ด้านในสุดอีกที
// เพิ่มเติมอีกตัวคือ ในการจัดการกับ Use case ต้องมีการจัดการผ่าน Repository (เป็นตัวแทนของการคุยกับ Database) ด้วย ดังนั้น ในการทำ Use case ต้องมีการเพิ่ม interface ของ Repository ด้วย เพื่อเป็นการบอกไปยัง Adapter ที่จะ implement - ว่าจะต้องส่งคำสั่งไหนมาบ้างเพื่อให้ใช้งานตาม Use case ได้
// ดังนั้น file ที่เกี่ยวข้องจะมี 2 files คือ

// order_repository.go เป็นตัวแทนของ interface ของ Repository
// order_use_case.go เป็นไฟล์สำหรับการเก็บ Business Logic ของ use case ไว้
type OrderRepository interface {
	Save(order entities.Order) error
}

// อธิบาย code

// order_repository.go ทำการเก็บ Repository ของ Order ไว้ จึงมี code เพียงแค่ interface ของ OrderRepository เท่านั้นและเป็นการบอกว่า Repository ต้องมี method Save() มาด้วย จึงจะตรงตาม OrderRepository
// order_use_case.go ทำการเก็บคำสั่งสำหรับจัดการสร้าง order เอาไว้ซึ่งก็คือ CreateOrder() โดยจะทำการ implement ผ่าน method ของ usecase ไว้ในตัวชื่อ OrderUseCase และจะมี OrderService เป็นตัวแทนของการรับ OrderRepository จากภายนอก (ที่ส่งเข้ามาจาก Adapter) มาอีกที
