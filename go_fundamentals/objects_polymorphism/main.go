package main

import (
	"bufio"
	"fmt"
	"objects_poly/menu"
	"os"
	"strings"
)

func main() {

loop:
	for {
		fmt.Println("Select an option")
		fmt.Println("1) Print Menu")
		fmt.Println("2) Add Item")
		fmt.Println("q) Quit")
		var in = bufio.NewReader(os.Stdin)
		choice, _ := in.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			menu.PrintMenu()
		case "2":
			menu.AddItem()
		case "q":
			fmt.Println("Quit")
			break loop
		default:
			break loop
		}
	}

}
