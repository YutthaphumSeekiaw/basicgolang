package main

import (
	"basic/error"
	"basic/function"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	a := 10
	id := uuid.New()
	fmt.Printf("#%v\n", &a)
	fmt.Println("Hello world")
	fmt.Printf("generate uuid : %s\n", id)
	sayHello()
	sayHello()
	function.HelloWorld()
	function.SayHello("Yutthaphum")

	number1 := 3
	number2 := 5
	sumNumber := function.Add(number1, number2)
	fmt.Println(sumNumber)
	error.LoginUser()

}

func sayHello() {
	fmt.Printf("Say Hello\n")
}

// receiver method (นิยาม method ใน go)
// ทีนี้สิ่งแรกที่เราต้องรู้ (สำหรับคนที่เรียน OOP หรือใช้ภาษา OOP อย่าง Java มา) นั่นคือ

// "Go ไม่มี Class" = เราจะไม่มี method ที่สามารถเข้าถึง element ใน object แบบเจาะจงได้

// นึกถึงตอนใช้ Class ปกติเราจะมีการสร้าง Class ต้นแบบไว้หนึ่งอัน
// เสร็จแล้วเราสร้าง Object โดยการอ้างอิงถึง Class ต้นแบบและนำมาใช้สำหรับข้อมูล Object ตัวนั้นออกมา
// เราลองนึกถึงเคสง่ายๆ เช่นเรามี Student ที่เราเก็บ ชื่อจริง (firstname) และ นามสกุล (lastname) เอาไว้

// แล้วเราอยาก implement ตัว function getFullname คู่เข้าไปกับตัวนั้น เพื่อให้สามารถแสดง fullname ข้อมูลชุดนั้นมาได้
// ถ้าเป็นเคสของ javascript เราจะสามารถแบบนี้ได้

// const student = {
//   firstname: 'boo',
//   lastname: 'seekiaw',
//   getFullname: () => `${this.firstname} ${this.lastname}`
// }

// student.getFullname()

// คำถามคือ แล้วใน go ละ ? เราสามารถทำสิ่งนี้ได้ยังไง = คำตอบก็คือ เราต้องทำเป็น method ขึ้นมา

// method คือ function "ที่มี scope"
// ใน go จะมองว่า function คือกลุ่มของ code ที่รวมชุดคำสั่ง
// แต่จะมอง method เป็นเหมือน function ที่มี scope ของ object นั้น
// ซึ่งการทำให้ function กลายเป็น method ได้ จะต้องมีการเพิ่ม "receiver argument" เข้าไปหน้า function เพื่อเป็นการระบุ type ของ method และเราสามารถ instances method นั้นใหม่ "เสมือนว่า" เป็นการ new Object ในภาษา OOP ได้เลย

// โดยคุณสมบัติของ method (เมื่อมีการประกาศใช้เป็น method นั้น) จะมีดังนี้

// Tied to a Type = method จะถูก defined ผูกมัดไปกับ type ที่ประกาศไว้ใน method เสมอ เมื่อมีการ instances ตัวแปรใหม่ขึ้นมา
// Receiver Argument = เป็นเพียงการระบุถึงประเภทของ instance ที่ method เรียกถึงเท่านั้น
// สามารถ access เข้าถึง property ของ receiver ได้ (เหมือนกับการ access object ใน oop)

// Define the Student struct
type Student struct {
	Firstname string
	Lastname  string
}

// Method with a receiver of type Student
// This method returns the full name of the student
// เอา type มาผูกไว้ข้าหน้า function fullname จาก  function จะกลายเป็น property นึงใน student ที่สามารถเรียกใช้ได้
func (s Student) FullName() string {
	return s.Firstname + " " + s.Lastname
}

func ConcatFullName() {
	student := Student{
		Firstname: "Mike",
		Lastname:  "Lopster",
	}

	// Call the FullName method on the Student instance
	fullName := student.FullName()
	fmt.Println("Full Name of the student:", fullName)

}

// Define a struct type
type Rectangle struct {
	Length float64
	Width  float64
}

// Method with a receiver of type Rectangle
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func GetArea() {
	rect := Rectangle{Length: 10, Width: 5}

	// Call the Area method on Rectangle instance
	area := rect.Area()
	fmt.Println("Area of rectangle:", area)
}
