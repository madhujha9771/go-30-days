package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("What is your name:")
	//var name string

	//fmt.Scan(&name)   //This will break on space
	//fmt.Println("Entered name is:", name)

	reader := bufio.NewReader(os.Stdin) // For taking user input with spaces we uses Bufio library in go
	name, _ := reader.ReadString('\n')

	fmt.Println("Entered name is:", name)

}
