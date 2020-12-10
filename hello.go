package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitorings = 2
const delay = 3

func main() {
	showIntro()

	for {
		showMenu()

		command := readCommand()

		switch command {
		case 1:
			initMonitoring()
		case 2:
			fmt.Println("Show logs ...")
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Command not found")
			os.Exit(-1)
		}
	}
}

func showIntro() {
	var name string
	fmt.Println("Hi, what's your name?")
	fmt.Scan(&name)
	fmt.Println("---------------------")
	fmt.Println("Hi,", name, "Select one action:")
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

func initMonitoring() {
	fmt.Println("Monitoring ...")

	sites := []string{"http://leandromachado.me/", "http://www.pasqualisolution.com.br/"}

	for i := 0; i < monitorings; i++ {
		for _, site := range sites {
			testSite(site)
			time.Sleep(delay * time.Second)
		}
		fmt.Println("")
	}
	fmt.Println("Successfully monitored :)")
	fmt.Println("")
}

func testSite(site string) {
	res, _ := http.Get(site)
	if res.StatusCode == 200 {
		fmt.Println(site, "it's work!")
	} else {
		fmt.Println(site, "not work :(")
	}
}
