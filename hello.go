package main

import (
	"fmt"
	"os"
)

func main() {
	showIntro()
	showMenu()

	command := readCommand()

	switch command {
	case 1:
		fmt.Println("Monitoring ...")
	case 2:
		fmt.Println("Show logs ...")
	case 0:
		os.Exit(0)
	default:
		fmt.Println("Command not found")
		os.Exit(-1)
	}
}

func showIntro() {
	var name string
	fmt.Println("Hi, what's your name?")
	fmt.Scan(&name)
	fmt.Println("---------------------")
	fmt.Println("Hi,", name)
}

func showMenu() {
	fmt.Println("1 - Init monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)

	return command
}
