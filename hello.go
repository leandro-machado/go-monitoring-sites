package main

import (
	"fmt"
)

func main() {
	name := "Leandro Machado"
	fmt.Println("Hi,", name)

	fmt.Println("1 - Init monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")

	var command int

	fmt.Scan(&command)

	fmt.Println("The command choosed is", command)
}
