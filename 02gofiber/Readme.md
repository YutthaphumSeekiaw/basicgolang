go fiber

# run แบบปกติ
go run  *.go
go run .

# run แบบ build
go build


https://gofiber.io/
go get github.com/gofiber/fiber/v2

https://docs.gofiber.io/guide/templates
go get github.com/gofiber/template/html/v2

Fiber คือ library ที่ได้แรงบันดาลใจมาจาก Express (ของฝั่ง node.js) ที่ build อยู่บน Fasthttp

Fasthttp คือ fastest HTTP engine for Go
เน้นไปที่ความไว และความสามารถในการจัดการ "zero memory allocation" ได้


# ENV
https://pkg.go.dev/github.com/joho/godotenv

go get github.com/joho/godotenv

# MiddleWare
https://docs.gofiber.io/category/-middleware

go get -u github.com/gofiber/jwt/v2

https://docs.gofiber.io/guide/grouping

Middleware คือ concept ของการเพิ่ม function คั่นกลางระหว่าง application เพื่อใช้สำหรับการเข้าถึง / modified request - response ของ request-response cycle เพื่อส่งต่อไปยัง function หลักต่อไปอีกที

ไอเดียใหญ่ๆของพวก Middleware คือ

1 ใช้สำหรับตรวจสอบ request / response เข้ามาก่อนได้ว่าถูกต้องหรือไม่ (เช่น เช็คว่า user login หรือไม่, user ถูก role หรือไม่)
2 ใช้สำหรับการเพิ่มข้อมูลบางอย่างคู่กับ request / response ไปเพื่อให้สามารถใช้กับ service function ที่เรียกคู่กันได้ (เช่น ส่งข้อมูล user เข้าไปคู่กับใน request เพื่อให้ทุก service function สามารถเรียกใช้งานได้)
3 เพื่อทำการเพิ่มเติมของบางอย่างระหว่างทางเข้าไป โดยใช้ข้อมูลจาก request / response นั้น (เช่น Log, Cache)


# Swagger
https://github.com/gofiber/swagger

go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/gofiber/swagger

package main

import (
  "fmt"
  "time"

  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/jwt/v2"
  "github.com/golang-jwt/jwt/v4"
  "github.com/gofiber/swagger"
  _ "github.com/mikelopster/fiber-basic/docs" // load generated docs
)

// @title Book API
// @description This is a sample server for a book API.
// @version 1.0
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
  // ... code เหมือนเดิม
}

// Handler functions
// getBooks godoc
// @Summary Get all books
// @Description Get details of all books
// @Tags books
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} Book
// @Router /book [get]
func getBooks(c *fiber.Ctx) error {
  // Retrieve user data from the context
  user := c.Locals(userContextKey).(*UserData)

  // Use the user data (e.g., for authorization, custom responses, etc.)
  fmt.Printf("User Email: %s, Role: %s\n", user.Email, user.Role)

  return c.JSON(books)
}


ทุกครั้งที่แก้เสร็จ ให้ run

# *******swag init

อีก 1 ที = จะได้ document ตัวใหม่ออกมา

# Logging
Logging: https://docs.gofiber.io/api/log