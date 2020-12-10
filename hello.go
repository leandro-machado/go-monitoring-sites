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

	if command == 1 {
		fmt.Println("Monitoring ...")
	} else if command == 2 {
		fmt.Println("Show logs ...")
	} else if command == 0 {
		fmt.Println("Exiting")
	}
}
