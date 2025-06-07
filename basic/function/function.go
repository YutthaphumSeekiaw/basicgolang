package function

import "fmt"

func HelloWorld() {
	fmt.Printf("Hello World\n")
}

func SayHello(name string) {
	fmt.Printf("Hello %s\n", name)
}

func Add(a int, b int) int {
	return a + b
}
