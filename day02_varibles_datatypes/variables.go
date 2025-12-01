package main

import "fmt"

// func main() {
	var name = "Shivam"
	var roll = 10
	var marks = 78.9
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Type of name %T\n", name)
	fmt.Printf("Roll no: go%d\n", roll)
	fmt.Printf("Type of roll %T\n", roll)
	fmt.Printf("Marks: %f\n", marks)
	fmt.Printf("Type of Marks %T\n", marks)
	var college string
	var fee float32
	var room_no int
	fmt.Println("Enter college name:")
	fmt.Scanf("%s\n", &name)
	fmt.Println("Fees of College:")
	fmt.Scanf("%f\n", &fee)
	fmt.Println("Enter room_no: ")
	fmt.Scanf("%d\n", &room_no)
	fmt.Println("College Name: %s", college)
	fmt.Println("Fees of collage: %f", fee)
	fmt.Println("Room_no : %d", room_no)
}
