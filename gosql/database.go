package main

type Product struct {
	// define your product fields
	ID         int
	Name       string
	Price      int
	Category   string
	Quantity   int
	SupplierID int
}

type Supplier struct {
	// define your supplier fields
	Name     string
	Location string
}
