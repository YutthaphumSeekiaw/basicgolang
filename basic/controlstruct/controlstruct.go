package controlstruct

import "fmt"

func Control() {
	var score int = 62 // ตัวอย่างสมมุติว่า 62 คะแนน

	if score >= 70 {
		fmt.Printf("PASSED")
	} else {
		// ก็จะทำงานตรงตำแหน่งนี้ เนื่องจากเงื่อนไขอันบนเป็นเท็จ
		fmt.Printf("FAILED")
	}
}

func Control2() {
	var score int = 62
	var grade string

	if score >= 80 {
		grade = "A"
	} else if score >= 70 {
		grade = "B"
	} else if score >= 60 {
		grade = "C"
	} else {
		grade = "F"
	}

	fmt.Printf("Your grade is %s", grade)
}

func Control3() {
	var dayOfWeek = 3

	switch dayOfWeek {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid Day")
	}
}

func Control4() {
	var score int = 62
	var grade string

	switch {
	case score >= 80:
		grade = "A"
	case score >= 70:
		grade = "B"
	case score >= 60:
		grade = "C"
	default:
		grade = "F"
	}

	fmt.Printf("Your grade is %s", grade)
}

func Control5() {
	num1 := 5
	num2 := 10

	sumNum := num1 + num2

	if sumNum >= 10 {
		fmt.Println("sumNum more than 10")
	}
}

func Control6() {
	num1 := 5
	num2 := 10

	if sumNum := num1 + num2; sumNum >= 10 {
		fmt.Println("sumNum more than 10")
	}
}

func ForLoop() {
	for i := 1; i < 10; i++ {
		fmt.Printf("number: %d", i)
	}
}

func DoWhile() {
	i := 1
	for {
		fmt.Printf("number: %d\n", i)
		i++
		if i >= 10 {
			break
		}
	}
}

func While() {
	i := 1
	for i < 10 {
		fmt.Printf("number: %d\n", i)
		i++
	}
}
