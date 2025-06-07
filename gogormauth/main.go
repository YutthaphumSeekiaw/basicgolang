// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// const (
// 	host     = "localhost"  // or the Docker service name if running in another container
// 	port     = 5432         // default PostgreSQL port
// 	user     = "boo"        // as defined in docker-compose.yml
// 	password = "P@ssw0rd"   // as defined in docker-compose.yml
// 	dbname   = "pgdatabase" // as defined in docker-compose.yml
// )
//===============================     Example 1
// // func main() {
// // 	// Configure your PostgreSQL database details here
// // 	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
// // 		"password=%s dbname=%s sslmode=disable",
// // 		host, port, user, password, dbname)
// // 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// // 	if err != nil {
// // 		panic("failed to connect to database")
// // 	}
// // 	// Migrate the schema
// //     // สร้าง database ให้อัตโนมัต ตาม structs
// // 	db.AutoMigrate(&Book{})
// // 	fmt.Println("Database migration completed!")
// // }

//====================================================================================================     Example 2

// // เพิ่ม Tracing

// // //https://gorm.io/docs/logger.html
// // เพื่อให้สะดวกต่อการ debug ใน GORM เรื่อง query
// // GORM ได้ทำการเพิ่ม logger มาให้เพื่อให้สะดวกต่อการ debug มากขึ้น

// func main() {
// 	// Configure your PostgreSQL database details here
// 	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)

// 	// New logger for detailed SQL logging
// 	newLogger := logger.New(
// 		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
// 		logger.Config{
// 			SlowThreshold: time.Second, // Slow SQL threshold
// 			LogLevel:      logger.Info, // Log level
// 			Colorful:      true,        // Enable color
// 		},
// 	)

// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
// 		Logger: newLogger, // add Logger
// 	})

// 	if err != nil {
// 		panic("failed to connect to database")
// 	}

// 	// Migrate the schema
// 	db.AutoMigrate(&Book{})
// 	fmt.Println("Database migration completed!")
// }

//=====================================================================================     Example 3 กับ fiber

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// const (
// 	host     = "localhost"  // or the Docker service name if running in another container
// 	port     = 5432         // default PostgreSQL port
// 	user     = "boo"        // as defined in docker-compose.yml
// 	password = "P@ssw0rd"   // as defined in docker-compose.yml
// 	dbname   = "pgdatabase" // as defined in docker-compose.yml
// )

// func main() {
// 	// Configure your PostgreSQL database details here
// 	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)

// 	// New logger for detailed SQL logging
// 	newLogger := logger.New(
// 		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
// 		logger.Config{
// 			SlowThreshold: time.Second, // Slow SQL threshold
// 			LogLevel:      logger.Info, // Log level
// 			Colorful:      true,        // Enable color
// 		},
// 	)

// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
// 		Logger: newLogger,
// 	})

// 	if err != nil {
// 		panic("failed to connect to database")
// 	}

// 	// Migrate the schema
// 	db.AutoMigrate(&Book{})

// 	// Setup Fiber
// 	app := fiber.New()

// 	// CRUD routes
// 	app.Get("/books", func(c *fiber.Ctx) error {
// 		return getBooks(db, c)
// 	})
// 	app.Get("/books/:id", func(c *fiber.Ctx) error {
// 		return getBook(db, c)
// 	})
// 	app.Post("/books", func(c *fiber.Ctx) error {
// 		return createBook(db, c)
// 	})
// 	app.Put("/books/:id", func(c *fiber.Ctx) error {
// 		return updateBook(db, c)
// 	})
// 	app.Delete("/books/:id", func(c *fiber.Ctx) error {
// 		return deleteBook(db, c)
// 	})

// 	// Start server
// 	log.Fatal(app.Listen(":8000"))
// }

//==================================================================================== Example 3 fiber and midleware ทำ AUTH

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/golang-jwt/jwt/v4"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// const (
// 	host     = "localhost"  // or the Docker service name if running in another container
// 	port     = 5432         // default PostgreSQL port
// 	user     = "boo"        // as defined in docker-compose.yml
// 	password = "P@ssw0rd"   // as defined in docker-compose.yml
// 	dbname   = "pgdatabase" // as defined in docker-compose.yml
// )

// func authRequired(c *fiber.Ctx) error {
// 	cookie := c.Cookies("jwt")
// 	jwtSecretKey := "testsecret" // should be env

// 	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return jwtSecretKey, nil
// 	})

// 	if err != nil || !token.Valid {
// 		return c.SendStatus(fiber.StatusUnauthorized)
// 	}

// 	return c.Next()
// }

// func main() {
// 	// Code GORM เหมือนเดิม
// 	// Configure your PostgreSQL database details here
// 	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)

// 	// New logger for detailed SQL logging
// 	newLogger := logger.New(
// 		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
// 		logger.Config{
// 			SlowThreshold: time.Second, // Slow SQL threshold
// 			LogLevel:      logger.Info, // Log level
// 			Colorful:      true,        // Enable color
// 		},
// 	)

// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
// 		Logger: newLogger, // add Logger
// 	})

// 	if err != nil {
// 		panic("failed to connect to database")
// 	}

// 	// Migrate the schema
// 	db.AutoMigrate(&Book{})
// 	fmt.Println("Database migration completed!")

// 	// Setup Fiber
// 	app := fiber.New()

// 	// add For prevent auth
// 	app.Use("/books", authRequired)

// 	// CRUD routes = API Set book ตัวเดิม

// 	app.Post("/register", func(c *fiber.Ctx) error {
// 		return createUser(db, c)
// 	})

// 	app.Post("/login", func(c *fiber.Ctx) error {
// 		return loginUser(db, c)
// 	})

// 	// Start server
// 	log.Fatal(app.Listen(":8000"))
// }

// ========================================================================== Example 4 relation and gorm

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	// New logger for detailed SQL logging
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger, // add Logger
	})

	if err != nil {
		panic("failed to connect to database")
	}

	// Migrate the schema
	db.AutoMigrate(&Book{}, &Publisher{}, &Author{}, &AuthorBook{})

	// ขาสร้าง
	publisher := Publisher{
		Details: "Publisher Details",
		Name:    "Publisher Name",
	}
	_ = createPublisher(db, &publisher)

	// Example data for a new author
	author := Author{
		Name: "Author Name",
	}
	_ = createAuthor(db, &author)

	// // Example data for a new book with an author
	book := Book{
		Name:        "Book Title",
		Author:      "Book Author",
		Description: "Book Description",
		PublisherID: publisher.ID,     // Use the ID of the publisher created above
		Authors:     []Author{author}, // Add the created author
	}
	_ = createBookWithAuthor(db, &book, []uint{author.ID})

	// ขาเรียก

	// Example: Get a book with its publisher
	bookWithPublisher, err := getBookWithPublisher(db, 1) // assuming a book with ID 1
	if err != nil {
		// Handle error
	}

	// Example: Get a book with its authors
	bookWithAuthors, err := getBookWithAuthors(db, 1) // assuming a book with ID 1
	if err != nil {
		// Handle error
	}

	// Example: List books of a specific author
	authorBooks, err := listBooksOfAuthor(db, 1) // assuming an author with ID 1
	if err != nil {
		// Handle error
	}

	fmt.Println(bookWithPublisher)
	fmt.Println(bookWithAuthors)
	fmt.Println(authorBooks)
}
