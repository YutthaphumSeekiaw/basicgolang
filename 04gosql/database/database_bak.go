package database

// import (
// 	"database/sql"
// )

// // Note

// // Exec = เป็น method ที่ใช้สำหรับการ execute SQL โดยไม่มีการ return rows กลับคืนมา (เช่น INSERT, UPDATE, DELETE)
// // QueryRow = เป็น method ที่ใช้สำหรับ query SQL เพื่อดึงข้อมูลกลับมา (เป็นข้อมูลตัวเดียว) ปกติจะใช้กับตระกูลของ SELECT

// type Product struct {
// 	// define your product fields
// 	ID         uint
// 	Name       string
// 	Price      int
// 	Category   string
// 	Quantity   int
// 	SupplierID int
// }

// type Supplier struct {
// 	// define your supplier fields
// 	Name     string
// 	Location string
// }

// func createProduct(db *sql.DB, name string, price int, category string, quantity int) (int, error) {
// 	var id int
// 	err := db.QueryRow(`INSERT INTO products(name, price, category, quantity) VALUES($1, $2, $3, $4) RETURNING id;`, name, price, category, quantity).Scan(&id)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return id, nil
// }

// func getProduct(db *sql.DB, id int) (Product, error) {
// 	var p Product
// 	row := db.QueryRow(`SELECT id, name, price, category, quantity FROM products WHERE id = $1;`, id)
// 	err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Category, &p.Quantity)
// 	if err != nil {
// 		return Product{}, err
// 	}
// 	return p, nil
// }

// func updateProduct(db *sql.DB, id int, name string, price int, category string, quantity int) error {
// 	_, err := db.Exec(`UPDATE products SET name = $1, price = $2, category = $3, quantity = $4 WHERE id = $5;`, name, price, category, quantity, id)
// 	return err
// }

// func deleteProduct(db *sql.DB, id int) error {
// 	_, err := db.Exec(`DELETE FROM products WHERE id = $1;`, id)
// 	return err
// }

// func getProducts(db *sql.DB) ([]Product, error) {
// 	rows, err := db.Query("SELECT id, name, price, category, quantity FROM products")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	// เพิ่มเติม
// 	// defer คือ คำสั่งที่จะทำงานก่อนคำสั่งสุดท้ายใน program (หรือ function นั้นๆ) โดยปกติมันจะใช้สำหรับคำสั่ง cleanup เพื่อปิด process ให้ครบก่อนที่จะหยุดทำงาน

// 	var products []Product
// 	for rows.Next() {
// 		var p Product
// 		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Category, &p.Quantity)
// 		if err != nil {
// 			return nil, err
// 		}
// 		products = append(products, p)
// 	}

// 	// Check for errors from iterating over rows
// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return products, nil
// }

// type ProductWithSupplier struct {
// 	ProductID        int
// 	ProductName      string
// 	Price            int
// 	SupplierName     string
// 	SupplierLocation string
// }

// func getProductsAndSuppliers(db *sql.DB) ([]ProductWithSupplier, error) {
// 	// SQL JOIN query
// 	query := `
//       SELECT
//           p.id AS product_id,
//           p.name AS product_name,
//           p.price,
//           s.name AS supplier_name,
//           s.location AS supplier_location
//       FROM
//           products p
//       INNER JOIN suppliers s
//           ON p.supplier_id = s.id;`

// 	rows, err := db.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var products []ProductWithSupplier
// 	for rows.Next() {
// 		var p ProductWithSupplier
// 		err := rows.Scan(&p.ProductID, &p.ProductName, &p.Price, &p.SupplierName, &p.SupplierLocation)
// 		if err != nil {
// 			return nil, err
// 		}
// 		products = append(products, p)
// 	}

// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return products, nil
// }

// // Transaction
// // Docs : https://go.dev/doc/database/execute-transactions
// func addProductAndSupplier(db *sql.DB, product Product, supplier Supplier) error {
// 	// Start a transaction
// 	tx, err := db.Begin()
// 	if err != nil {
// 		return err
// 	}

// 	// Rollback the transaction in case of a panic
// 	defer tx.Rollback()

// 	// Insert into the supplier table
// 	supplierResult, err := tx.Exec("INSERT INTO suppliers (name, location) VALUES ($1, $2) RETURNING id", supplier.Name, supplier.Location)
// 	if err != nil {
// 		return err
// 	}

// 	// Get the last inserted ID for the supplier
// 	supplierID, err := supplierResult.LastInsertId()
// 	if err != nil {
// 		return err
// 	}

// 	// Insert into the product table
// 	_, err = tx.Exec("INSERT INTO products (name, price, category, quantity, supplier_id) VALUES ($1, $2, $3, $4, $5)", product.Name, product.Price, product.Category, product.Quantity, supplierID)
// 	if err != nil {
// 		return err
// 	}

// 	// Commit the transaction
// 	return tx.Commit()
// }

// // Note

// // Transaction ใน go คือ เปิดผ่านคำสั่ง db.Begin()
// // เมื่อ run ทุกอย่างเรียบร้อยให้ใช้คำสั่ง tx.Commit()
// // และถ้ามีอะไรเกิดขึ้นไม่ถูกต้อง ใช้คำสั่ง tx.Rollback()
// // โดยปกติการ handle Rollback จะใช้ defer tx.Rollback() เพื่อเป็นการ run เสมอตอนจบ function (และคำสั่งจะไม่มีผลอะไรหาก run tx.Commit() ไปแล้วเรียบร้อย)
