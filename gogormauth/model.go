// package main

// import (
// 	"fmt"
// 	"log"

// 	"gorm.io/gorm"
// )

// type Book struct {
// 	gorm.Model
// 	Name        string
// 	Author      string
// 	Description string
// }

// func CreateBook(db *gorm.DB, book *Book) {
// 	result := db.Create(book)
// 	if result.Error != nil {
// 		log.Fatalf("Error creating book: %v", result.Error)
// 	}
// 	fmt.Println("Book created successfully")
// }

// func GetBook(db *gorm.DB, id uint) *Book {
// 	var book Book
// 	result := db.First(&book, id)
// 	if result.Error != nil {
// 		log.Fatalf("Error finding book: %v", result.Error)
// 	}
// 	return &book
// }

// func UpdateBook(db *gorm.DB, book *Book) {
// 	result := db.Save(book)
// 	if result.Error != nil {
// 		log.Fatalf("Error updating book: %v", result.Error)
// 	}
// 	fmt.Println("Book updated successfully")
// }

// func DeleteBook(db *gorm.DB, id uint) {
// 	var book Book
// 	result := db.Delete(&book, id)
// 	if result.Error != nil {
// 		log.Fatalf("Error deleting book: %v", result.Error)
// 	}
// 	fmt.Println("Book deleted successfully")
// }

// // การ Delete ใน GORM
// // Ref: https://gorm.io/docs/delete.html

// // Default ของ Gorm คือ Soft Delete
// // Soft Delete คือ feature ที่ record ไม่ได้โดนลบไปจริงๆ แต่แค่จะทำการ mark ว่า delete ไว้เฉยๆ
// // ใน GORM default จะทำการสร้าง field ที่ชื่อ DeletedAt ไว้ เมื่อมีการลบจะเป็นการ update วันที่ของ DeletedAt เอาไว้

// func HardDeleteBook(db *gorm.DB, id uint) error {
// 	var book Book
// 	result := db.Unscoped().Delete(&book, id)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	fmt.Println("Book hard deleted successfully")
// 	return nil
// }

// // Search Sort
// func getBooksSortedByCreatedAt(db *gorm.DB) ([]Book, error) {
// 	var books []Book
// 	result := db.Order("created_at desc").Find(&books)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return books, nil
// }

// func getBooksByAuthorName(db *gorm.DB, authorName string) ([]Book, error) {
// 	var books []Book
// 	result := db.Where("author = ?", authorName).Find(&books)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return books, nil
// }

//============================================================================= Example 3 fiber

// package main

// import (
//   "github.com/gofiber/fiber/v2"
//   "gorm.io/gorm"
// )

// type Book struct {
//   gorm.Model
//   Name        string `json:"name"`
//   Author      string `json:"author"`
//   Description string `json:"description"`
// }

// // getBooks retrieves all books
// func getBooks(db *gorm.DB, c *fiber.Ctx) error {
//   var books []Book
//   db.Find(&books)
//   return c.JSON(books)
// }

// // getBook retrieves a book by id
// func getBook(db *gorm.DB, c *fiber.Ctx) error {
//   id := c.Params("id")
//   var book Book
//   db.First(&book, id)
//   return c.JSON(book)
// }

// // createBook creates a new book
// func createBook(db *gorm.DB, c *fiber.Ctx) error {
//   book := new(Book)
//   if err := c.BodyParser(book); err != nil {
//     return err
//   }
//   db.Create(&book)
//   return c.JSON(book)
// }

// // updateBook updates a book by id
// func updateBook(db *gorm.DB, c *fiber.Ctx) error {
//   id := c.Params("id")
//   book := new(Book)
//   db.First(&book, id)
//   if err := c.BodyParser(book); err != nil {
//     return err
//   }
//   db.Save(&book)
//   return c.JSON(book)
// }

// // deleteBook deletes a book by id
// func deleteBook(db *gorm.DB, c *fiber.Ctx) error {
//   id := c.Params("id")
//   db.Delete(&Book{}, id)
//   return c.SendString("Book successfully deleted")
// }

//========================================================================= Example 3 fiber and midleware ทำ AUTH

// package main

// import (
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/golang-jwt/jwt/v4"
// 	"golang.org/x/crypto/bcrypt"
// 	"gorm.io/gorm"
// )

// type User struct {
// 	gorm.Model
// 	Email    string `gorm:"unique"`
// 	Password string
// }

// type Book struct {
// 	gorm.Model
// 	Name        string
// 	Author      string
// 	Description string
// }

// // createUser handles user registration
// func createUser(db *gorm.DB, c *fiber.Ctx) error {
// 	user := new(User)
// 	if err := c.BodyParser(user); err != nil {
// 		return err
// 	}

// 	// Encrypt the password
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return err
// 	}
// 	user.Password = string(hashedPassword)

// 	// Create user
// 	db.Create(user)
// 	return c.JSON(user)
// }

// // loginUser handles user login
// func loginUser(db *gorm.DB, c *fiber.Ctx) error {
// 	var input User
// 	var user User

// 	if err := c.BodyParser(&input); err != nil {
// 		return err
// 	}

// 	// Find user by email
// 	db.Where("email = ?", input.Email).First(&user)

// 	// Check password
// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
// 		return c.SendStatus(fiber.StatusUnauthorized)
// 	}

// 	// Create JWT token
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["user_id"] = user.ID
// 	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
// 	jwtSecretKey := "testsecret" // should be env

// 	t, err := token.SignedString(jwtSecretKey)
// 	if err != nil {
// 		return c.SendStatus(fiber.StatusInternalServerError)
// 	}

// 	// Set cookie
// 	c.Cookie(&fiber.Cookie{
// 		Name:     "jwt",
// 		Value:    t,
// 		Expires:  time.Now().Add(time.Hour * 72),
// 		HTTPOnly: true,
// 	})

// 	return c.JSON(fiber.Map{"message": "success"})
// }

// =========================================================================== Example 4 relation and gorm

package main

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	PublisherID uint
	Publisher   Publisher
	Authors     []Author `gorm:"many2many:author_books;"`
}

type Publisher struct {
	gorm.Model
	Details string
	Name    string
}

type Author struct {
	gorm.Model
	Name  string
	Books []Book `gorm:"many2many:author_books;"`
}

type AuthorBook struct {
	AuthorID uint
	Author   Author
	BookID   uint
	Book     Book
}

func createPublisher(db *gorm.DB, publisher *Publisher) error {
	result := db.Create(publisher)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func createAuthor(db *gorm.DB, author *Author) error {
	result := db.Create(author)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func createBookWithAuthor(db *gorm.DB, book *Book, authorIDs []uint) error {
	// First, create the book
	if err := db.Create(book).Error; err != nil {
		return err
	}

	return nil
}

func getBookWithPublisher(db *gorm.DB, bookID uint) (*Book, error) {
	var book Book
	result := db.Preload("Publisher").First(&book, bookID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

func getBookWithAuthors(db *gorm.DB, bookID uint) (*Book, error) {
	var book Book
	result := db.Preload("Authors").First(&book, bookID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

func listBooksOfAuthor(db *gorm.DB, authorID uint) ([]Book, error) {
	var books []Book
	result := db.Joins("JOIN author_books on author_books.book_id = books.id").
		Where("author_books.author_id = ?", authorID).
		Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}
