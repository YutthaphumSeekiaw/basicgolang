package datastruct

import "fmt"

func Array() {
	var myArray [3]int // An array of 3 integers
	myArray[0] = 10    // Assign values
	myArray[1] = 20
	myArray[2] = 30

	// Reassigning the elements of the array
	myArray[0] = 100

	myArray[1] = 200
	myArray[2] = 300

	// Looping through the array
	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}
	fmt.Println(myArray) // Output: [100 200 300]
}

func Slice() {
	// 	slice จะมีความแตกต่างกับ Array คือ
	// ไม่ต้องประกาศขนาด สามารถประกาศ [] แล้วใช้งานได้เลย
	// จะกำหนดค่าเริ่มต้น หรือไม่กำหนดค่าเริ่มต้นก็ได้
	// สามารถใช้คำสั่ง append ในการเพิ่มข้อมูลเข้า array ได้ (ด้วยคุณสมบัติที่ยืดหยุ่นทำให้เราสามารถใช้คำสั่ง append ทั้งเพิ่มและลบข้อมูลได้)
	// สามารถระบุ index ระหว่างกลางเพื่อดึงข้อมูลออกมาตามขนาดที่เพิ่มขึ้น เช่น myslice[1:3] เก็บการหยิบ index ตั้งแต่ 1 ถึง 3 ออกมา (Array ก็สามารถทำได้ แต่ต้องทำในขนาดที่ประกาศเอาไว้)

	mySlice := []int{10, 20, 30, 40, 50} // A slice of integers

	fmt.Println(mySlice)      // Output: [10 20 30 40 50]
	fmt.Println(len(mySlice)) // Length of the slice: 5
	fmt.Println(cap(mySlice)) // Capacity of the slice: 5

	// Slicing a slice
	subSlice := mySlice[1:3] // Slice from index 1 to 2
	fmt.Println(subSlice)    // Output: [20 30]

	var mySlice2 []int // Declared but not initialized

	// Appending data to the slice
	mySlice2 = append(mySlice2, 10)
	mySlice2 = append(mySlice2, 20, 30)

	fmt.Println(mySlice2) // Output: [10 20 30]
}

func ConvertArrayToSlice() {
	var myArray [3]int // An array of 3 integers
	myArray[0] = 10    // Assign values
	myArray[1] = 20
	myArray[2] = 30

	// Converting array to slice
	mySlice := myArray[:]

	// Resizing slice by appending new elements
	mySlice = append(mySlice, 40, 50)

	fmt.Println(mySlice) // Output will show a slice with 5 elements: [10 20 30 40 50]
}

func Map() {
	// Map คือ data type ที่ใช้สำหรับเก็บ key คู่กับ value เอาไว้

	// 	ตัวอย่างการใช้ map กัน

	// เริ่มต้นด้วยการประกาศ make(map[<ประเภท key>]<ประเภท value>)
	// คำสั่ง make คือคำสั่ง initialize maps เพื่อเป็นการจอง memory สำหรับการใช้ map
	//  และสร้าง data structure ประเภทนี้ขึ้นมา (แบบ reference types) = เป็นคำสั่งสำหรับทำ memory allocation
	//   เพื่อให้แน่ใจว่ามี memory อยู่จริงสำหรับการใช้งานตัวแปร

	myMap := make(map[string]int)

	// Add key-value pairs to the map
	myMap["apple"] = 5
	myMap["banana"] = 10
	myMap["orange"] = 8

	// Access and print a value for a key
	fmt.Println("Apples:", myMap["apple"])

	// Update the value for a key
	myMap["banana"] = 12

	// Delete a key-value pair
	delete(myMap, "orange")

	// Iterate over the map
	for key, value := range myMap {
		fmt.Printf("%s -> %d\n", key, value)
	}

	// Checking if a key exists
	val, ok := myMap["pear"]
	if ok {
		fmt.Println("Pear's value:", val)
	} else {
		fmt.Println("Pear not found inmap")
	}
}

// ============= Struct =====================
type Student struct {
	Name   string
	Height int
	Weight int
	Grade  string
}

func Struct() {
	// Struct คือ data type ที่ทำการรวม data หลายประเภทเข้ามาไว้ในตัวแปรตัวเดียวกัน
	// Create an instance of the Student struct
	// var student1 Student
	// student1.Name = "Mikelopster"
	// student1.Weight = 60
	// student1.Height = 180
	// student1.Grade = "F"

	// Create an instance of the Student struct
	// var student1 Student = Student{
	//   Name:   "Mikelopster",
	//   Weight: 60,
	//   Height: 180,
	//   Grade:  "F",
	// }
	student1 := Student{
		Name:   "Mikelopster",
		Weight: 60,
		Height: 180,
		Grade:  "F",
	}

	// Print struct values
	fmt.Println(student1)

}

func StructAndArray() {
	// Create an array of Student structs
	var students [3]Student

	// Initialize the first student
	students[0] = Student{
		Name:   "Mikelopster",
		Weight: 60,
		Height: 180,
		Grade:  "F",
	}

	// Initialize the second student
	students[1] = Student{
		Name:   "Alice",
		Weight: 55,
		Height: 165,
		Grade:  "A",
	}

	// Initialize the third student
	students[2] = Student{
		Name:   "Bob",
		Weight: 68,
		Height: 175,
		Grade:  "B",
	}

	// Print array of structs
	fmt.Println(students)
}

func StructAndMap() {
	// Create a map with string keys and Student struct values
	students := make(map[string]Student)

	// Add Student structs to the map
	students["st01"] = Student{Name: "Mikelopster", Weight: 60, Height: 180, Grade: "F"}
	students["st02"] = Student{Name: "Alice", Weight: 55, Height: 165, Grade: "A"}
	students["st03"] = Student{Name: "Bob", Weight: 68, Height: 175, Grade: "B"}

	// Print the map
	fmt.Println("Students Map:", students)

	// Access and print a specific student by key
	fmt.Println("Student st01:", students["st01"])
}

// Define a struct type
type Person struct {
	Name    string
	Age     int
	Address Address
}

// Another struct type used in Person
type Address struct {
	Street  string
	City    string
	ZipCode int
}

func StructAndStruct() {
	// Create an instance of the Person struct
	var person Person
	person.Name = "Alice"
	person.Age = 30
	person.Address = Address{
		Street:  "123 Main St",
		City:    "Gotham",
		ZipCode: 12345,
	}

	// Alternative way to initialize a struct
	bob := Person{
		Name: "Bob",
		Age:  25,
		Address: Address{
			Street:  "456 Elm St",
			City:    "Metropolis",
			ZipCode: 67890,
		},
	}

	// Print struct values
	fmt.Println(person)
	fmt.Println(bob)
}
