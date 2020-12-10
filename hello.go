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

	switch command {
	case 1:
		fmt.Println("Monitoring ...")
	case 2:
		fmt.Println("Show logs ...")
	case 0:
		fmt.Println("Exiting")
	default:
		fmt.Println("Command not found")
	}
}
