// มาลองประยุกต์ใช้ Pattern ต่างๆ

// Pubsub

// Pub/sub เป็นรูปแบบการสื่อสารระหว่าง goroutines ต่าง ๆ โดย goroutine หนึ่งสามารถส่งข้อความไปยัง
// goroutine อื่น ๆ หลายตัวได้ ในการประยุกต์ใช้ pub/sub กับ goroutines เราสามารถใช้ channel เพื่อส่งข้อความระหว่าง goroutines

//================================= ตัวอย่าง Pattern ทั่วไป
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	// สร้าง channel เพื่อส่งข้อความ
// 	ch := make(chan string)

// 	// สร้าง goroutine เพื่อส่งข้อความไปยัง channel
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			ch <- fmt.Sprintf("Hello, world! %d", i)
// 			time.Sleep(1 * time.Second)
// 		}
// 	}()

// 	// สร้าง goroutine เพื่อรับข้อความจาก channel
// 	go func() {
// 		for {
// 			msg := <-ch
// 			fmt.Println(msg)
// 		}
// 	}()

// 	// รอให้ goroutines ทำงานเสร็จสิ้น
// 	time.Sleep(5 * time.Second)
// }

//================================= ตัวอย่าง การใช้งานร่วมกับ Fiber

// Note

// ไอเดียคือ เราจะใช้ Fiber รับคำสั่งมา
// และนำคำสั่งนั้น ส่งผ่าน channel เข้าไป เพื่อให้ subscribe ข้อมูลเอาไว้ได้
// ลง Package Fiber

// go get github.com/gofiber/fiber/v2
// package main

// import (
// 	"fmt"

// 	"github.com/gofiber/fiber/v2"
// )

// var pubsub = &PubSub{
// 	Subscribers: make(map[uint]*Subscriber),
// }

// const (
// 	host     = "localhost"  // or the Docker service name if running in another container
// 	port     = 5432         // default PostgreSQL port
// 	user     = "boo"        // as defined in docker-compose.yml
// 	password = "P@ssw0rd"   // as defined in docker-compose.yml
// 	dbname   = "pgdatabase" // as defined in docker-compose.yml
// )

// func main() {
// 	app := fiber.New()

// 	// Start a subscription routine
// 	go func() {
// 		subscriber := pubsub.Subscribe()
// 		defer pubsub.Unsubscribe(subscriber.ID)

// 		for msg := range subscriber.Channel {
// 			// Handle the message, e.g., log, process, etc.
// 			fmt.Printf("Received message: %s\n", msg.Content)
// 		}
// 	}()

// 	// Publish endpoint
// 	app.Post("/publish", func(c *fiber.Ctx) error {
// 		var msg Message
// 		if err := c.BodyParser(&msg); err != nil {
// 			return err
// 		}

// 		go pubsub.Publish(msg)
// 		return c.SendString("Message published")
// 	})

// 	app.Listen(":8888")
// }

//================================= ตัวอย่าง  Cronjob

// Cronjob เป็นเครื่องมือที่ใช้ในการเรียกใช้งานงานซ้ำ ๆ ในช่วงเวลาที่กำหนด ในการประยุกต์ใช้ cronjob กับ goroutines
// เราสามารถใช้ goroutine เพื่อเรียกใช้งานงานซ้ำ ๆ ในช่วงเวลาที่กำหนด โดยเราสามารถกำหนดช่วงเวลาที่ต้องการเรียกใช้งานงานโดยใช้ channel
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	// สร้าง channel เพื่อกำหนดช่วงเวลาที่ต้องการเรียกใช้งานงาน
// 	ch := time.Tick(5 * time.Second)

// 	// สร้าง goroutine เพื่อเรียกใช้งานงานซ้ำ ๆ ในช่วงเวลาที่กำหนด
// 	go func() {
// 		for range ch {
// 			fmt.Println("Hello, world!")
// 		}
// 	}()

// 	// รอให้ goroutine ทำงานเสร็จสิ้น
// 	time.Sleep(60 * time.Second)
// }

//================================= ตัวอย่าง การใช้งานร่วมกับ gorm
// go get -u gorm.io/gorm
// go get -u gorm.io/driver/postgres

package main

import (
	"fmt"
	"log"

	"github.com/robfig/cron/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	user     = "boo"        // as defined in docker-compose.yml
	password = "P@ssw0rd"   // as defined in docker-compose.yml
	dbname   = "pgdatabase" // as defined in docker-compose.yml
)

func main() {
	// Configure your PostgreSQL database details here
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	// Migrate the schema
	db.AutoMigrate(&ExampleModel{})
	fmt.Println("Database migration completed!")

	c := cron.New()
	_, err = c.AddFunc("@every 1m", func() {
		go task(db)
	})

	if err != nil {
		log.Fatal("Error scheduling a task:", err)
	}

	c.Start()

	// Block the main thread as the cron job runs in the background
	select {}
}
