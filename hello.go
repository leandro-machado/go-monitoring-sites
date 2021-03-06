package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

const monitorings = 2
const delay = 1

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
			showLog()
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
	fmt.Println("Hi,", name)
	fmt.Println("Select one action:")
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
	res, err := http.Get(site)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	if res.StatusCode == 200 {
		fmt.Println(site, "it's work!")
		createLog(site, true)
	} else {
		fmt.Println(site, "not work :(")
	}
}

func createLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error:", file)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func showLog() {
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(string(file))
}
