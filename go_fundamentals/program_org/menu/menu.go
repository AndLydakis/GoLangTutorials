package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type MenuItem struct {
	name   string
	prices map[string]float64
}

var in = bufio.NewReader(os.Stdin)

func PrintMenu() {
	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, price)
		}
	}
}

func AddItem() {
	fmt.Println("Please enter new item")
	name, _ := in.ReadString('\n')
	menu = append(menu, MenuItem{name: name, prices: make(map[string]float64)})
	PrintMenu()
}
