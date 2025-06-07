# go mod init modulename

install package

# go get github.com/google/uuid

ทำการทดลอง run ภาษา go โดยใช้คำสั่ง

# go run main.go

ก็จะได้ผลลัพธ์เป็น "Hello World" ออกมาได้ ซึ่งในการพัฒนาจริงนั้น เราจะทำการ compile go ออกมาเป็นตัว machine code (binary)

# go build main.go

เสร็จแล้วจะได้เป็น file build ออกมา

ถ้าเป็น Mac / Linux จะออกมาเป็น script file สามารถ run ได้จาก
./main

และจะได้คำว่า "Hello World" ออกมา

ถ้าเป็น Windows จะออกมาเป็น exe file

# *** module นั้น เราจะมี 2 files ที่เป็น file หลักคือ

go.mod list dependency ทั้งหมดของ go ไว้
go.sum คือไฟล์ที่เก็บเวอร์ชั่นที่ลงเอาไว้ (ใข้เช็คได้ว่าเรา load มาถูก version แล้วหรือไม่)

# *** สำหรับเรื่องของ package

go ไม่มี concept class ทุกอย่างจัดการผ่าน function และ package
ใน 1 folder สามารถมีได้เพียง package เดียวเท่านั้น (หากตั้งชื่อ package ต่างกัน ใน folder เดียวกัน = จะเกิด error ออกมา)

# ***************** Data

# Array = รูปแบบการเก็บข้อมูลเป็น sequence เก็บหลายข้อมูลในตัวแปรเดียว กันไว้ เช่น var a [5]int
# Slice = คล้ายๆ Array แต่อนุญาตให้เปลี่ยนขนาดได้ (จากแต่เดิม Array ที่ต้องระบุขนาดเสมอ) เช่น var s[]int
# Map = ข้อมูล map ที่คล้ายๆ dictionary ที่สามารถเก็บ key คู่กับ value ตรงๆไว้ได้ (จากเดิมต้องระบุเป็นตำแหน่ง) เช่น map[string]int
# Struct = ตัวแปรที่ประกอบไปด้วยกลุ่มของ Variable ออกมาเป็นตัวแปรเดียว (ลักษณะคล้ายๆ Object ในภาษาอื่นๆ) โดยสามารถกำหนดชื่อ field และ type คู่กันไว้ได้ เช่น
        type Person struct {
            Name string
            Age  int
        }

# Function = ในภาษา go function ถูกจัดอยู่ในประเภทหนึ่งของ data type โดยสามารถทำการประกาศเป็นเหมือนตัวแปรตัวแปรหนึ่งขึ้นมาได้ เช่น var squareFunc func(int) int = square เป็นการบอกว่า สร้างตัวแปร function รับเป็น integer และคืนค่าเป็น integer (เดี๋ยวเรามาลงลึกใน function อีกที)
# Interface = data type ที่จะทำการระบุ set ของ method ที่จะต้องมีในกลุ่มที่จะใช้ data type ของตัวนี้ (เป็นเหมือนสร้างต้นแบบของ function เอาไว้)
            type Shape interface {
                Area() float64
                Perimeter() float64
            }

# Pointer = data type ที่ทำการเก็บ memory address ของ variable เอาไว้ เช่น var p *int
# Channel = data type ใช้สำหรับการ communication ระหว่าง goroutines



