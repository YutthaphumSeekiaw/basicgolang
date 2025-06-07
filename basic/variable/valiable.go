package variable

import "fmt"

var booleanVar bool = true
var intVar int = 10
var floatVar float64 = 3.14
var stringVar string = "Hello Go"

func Variable() {
	// golang การประกาษตัวแปลจะมีค่า default อยู่เสมอเช่น  int จะเป็น 0 string จะเป็น ""

	fmt.Println("Boolean:", booleanVar)
	fmt.Println("Integer:", intVar)
	fmt.Println("Float:", floatVar)
	fmt.Println("String:", stringVar)

	// การประกาศตัวเเปลเเบบสั้น ใช้ภายใน func
	booleanVar2 := true
	intVar2 := 0
	floatVar2 := 0.0
	stringVar2 := ""

	fmt.Println("Boolean:", booleanVar2)
	fmt.Println("Integer:", intVar2)
	fmt.Println("Float:", floatVar2)
	fmt.Println("String:", stringVar2)

	a := 10
	b := 3
	fmt.Println(a + b) // 13
	fmt.Println(a - b) // 7
	fmt.Println(a * b) // 30
	fmt.Println(a / b) // 3
	fmt.Println(a % b) // 1

	fmt.Println(a == b) // false
	fmt.Println(a != b) // true
	fmt.Println(a > b)  // true
	fmt.Println(a < b)  // false
	fmt.Println(a >= b) // true
	fmt.Println(a <= b) // false

	c := true
	d := false
	fmt.Println(c && d) // false
	fmt.Println(c || d) // true
	fmt.Println(!c)     // false

	aa := 5
	bb := 10
	cc := aa + bb // c ก็จะทำการเก็บ 15 เอาไว้
	fmt.Println(cc)
}
