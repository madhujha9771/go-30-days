package main

import "fmt"

func simple() {
	fmt.Println("Simple function")
}

func add(a, b int) int {
	return (a + b)
}
func main() {
	simple()
	ans := add(3, 4)
	fmt.Println("Number added:", ans)
}
