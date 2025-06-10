package pointer

import "fmt"

// Pointer คืออะไร ?
// Pointer คือ data type อีก 1 ประเภทที่ใช้สำหรับเก็บ memory address ของตัวแปร ถูกใช้ใช้ใน 2 จุดประสงค์ใหญ่ๆคือ

// Mutability = อนุญาตให้สามารถกลับมาแก้ original data ของตัวแปรตัวนั้นได้ (เพื่อให้เป็นการส่งแบบ pass by reference แทน)
// Efficient large structs = สามารถทำให้ส่งข้อมูลขนาดใหญ่ต่อไปได้ (โดยการส่ง address แทน value ของตัวแปรแทน) เช่น
// การส่ง struct ขนาดใหญ่เข้าไป
// การส่ง config ขนาดใหญ่เข้าไป (เช่น config database, connection ต่างๆ)

func Pointer1() {
	// Declare an integer variable
	x := 10

	// Declare a pointer to an integer and assign it the address of x
	var p *int = &x

	// Print the value of x and the value at the pointer p
	fmt.Println("Value of x:", x)  // Output: Value of x: 10
	fmt.Println("Value at p:", *p) // Output: Value at p: 10

	// Modify the value at the pointer p
	*p = 20

	// x is modified since p points to x
	fmt.Println("New value of x:", x) // Output: New value of x: 20
}

// ************** ในภาษา go ถ้าส่งค่าเป็น value ไป จะเเก้ไขข้อมูลไม่ได้ ต้องส่ง address pointer เพื่อเปลี่ยนค่าได้

// func changeValue(val int) {
// 	val = 50
// }

// func main() {
// 	x := 20
// 	changeValue(x)
// 	fmt.Println(x) // Output: 20 (x is unchanged)
// }

func changeValue(ptr *int) {
	*ptr = 50
}

func SetValue() {
	x := 20
	changeValue(&x)
	fmt.Println(x) // Output: 50 (x is changed)
}

// ======================
// type Person struct {
// 	Name string
// }

// func changeName(p Person) {
// 	p.Name = "Alice"
// }

// func main() {
// 	person := Person{Name: "Bob"}
// 	changeName(person)
// 	fmt.Println(person.Name) // Output: Bob (person.Name is unchanged)
// }

type Employee struct {
	Name   string
	Salary int
}

// Function to give a raise to an employee
func giveRaise(e *Employee, raise int) {
	e.Salary += raise
}

func SetEmpSalary() {
	emp := Employee{Name: "John Doe", Salary: 50000}

	giveRaise(&emp, 5000)
	fmt.Println("After raise:", emp)
}

// ==================
type ListNode struct {
	Value int
	Next  *ListNode
}

// Function to add a node to the front of the list
func prepend(head **ListNode, value int) {
	newNode := ListNode{Value: value, Next: *head}
	*head = &newNode
}

func LinkList() {
	var head *ListNode

	prepend(&head, 10)
	prepend(&head, 20)

	current := head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}

// ======================
// Config represents the application configuration
type Config struct {
	LogLevel string
	Port     int
}

// UpdateConfig modifies the provided configuration
func UpdateConfig(c *Config, logLevel string, port int) {
	c.LogLevel = logLevel
	c.Port = port
}

func GetConfig() {
	// Initial configuration
	appConfig := &Config{
		LogLevel: "info",
		Port:     8080,
	}

	fmt.Println("Initial Config:", appConfig)

	// Update configuration
	UpdateConfig(appConfig, "debug", 9000)
	fmt.Println("Updated Config:", appConfig)
}
