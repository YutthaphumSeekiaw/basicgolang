package main

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq" //เป็นการ import  เพื่อ lib ให้ตัวอื่น ใช้ method ได้เฉยๆ
)

// const (
// 	host     = "localhost"  // or the Docker service name if running in another container
// 	port     = 5432         // default PostgreSQL port
// 	user     = "boo"        // as defined in docker-compose.yml
// 	password = "P@ssw0rd"   // as defined in docker-compose.yml
// 	dbname   = "pgdatabase" // as defined in docker-compose.yml
// )

// func main() {
// 	// Connection string
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)

// 	// Open a connection
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Check the connection
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Successfully connected!")
// }

// ตัวอย่างการใช้ defer
// func main() {
//   fmt.Println(test())
//   fmt.Println("Successfully connected!")
// }

// func test() string {
//   defer fmt.Println("Before!")
//   fmt.Println("After!")
//   return "test"
// }

// ประกาศตัวแปร global สำหรับเรียกใช้ database
var db *sql.DB

func SetupDatabase() *sql.DB {
	const connectionString = "user=postgres password=yourpassword dbname=yourdbname sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

// ประกาศ struct Product สำหรับการรับ request
// type Product struct {
//   ID       int    `json:"id"`
//   Name     string `json:"name"`
//   Price    int    `json:"price"`
//   Category string `json:"category"`
// }

func main() {
	app := fiber.New()
	db = SetupDatabase()
	defer db.Close()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	// Set up routes
	app.Post("/products", CreateProduct)
	app.Get("/products/:id", GetProduct)
	app.Put("/products/:id", UpdateProduct)
	app.Delete("/products/:id", DeleteProduct)

	log.Fatal(app.Listen(":3000"))
}

func CreateProduct(c *fiber.Ctx) error {
	p := new(Product)
	if err := c.BodyParser(p); err != nil {
		return err
	}

	// Insert product into database
	_, err := db.Exec("INSERT INTO products (name, price, category) VALUES ($1, $2, $3)", p.Name, p.Price, p.Category)
	if err != nil {
		return err
	}

	return c.JSON(p)
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var p Product

	// Retrieve product from database
	err := db.QueryRow("SELECT id, name, price, category FROM products WHERE id = $1", id).Scan(&p.ID, &p.Name, &p.Price, &p.Category)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).SendString("Product not found")
		}
		return err
	}

	return c.JSON(&p)
}

func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	p := new(Product)
	if err := c.BodyParser(p); err != nil {
		return err
	}

	// Update product in the database
	_, err := db.Exec("UPDATE products SET name = $1, price = $2, category = $3 WHERE id = $4", p.Name, p.Price, p.Category, id)
	if err != nil {
		return err
	}

	p.ID, _ = strconv.Atoi(id)
	return c.JSON(p)
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	// Delete product from database
	_, err := db.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
