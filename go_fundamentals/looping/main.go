package main

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

func main() {
	fmt.Println("Select an option")
	fmt.Println("1) Print Menu")
	in := bufio.NewReader(os.Stdin)
	choice, _ := in.ReadString('\n')
	choice = strings.TrimSpace(choice)
	fmt.Println(choice)

	menu := []MenuItem{
		{name: "Coffee", prices: map[string]float64{"small": 1.65, "medium": 2.5, "large": 3.0}},
		{name: "Espresso", prices: map[string]float64{"small": 1.65, "medium": 2.5, "large": 3.0}},
	}
	fmt.Println(menu)

	for _, v := range menu {
		fmt.Println(v.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range v.prices {
			fmt.Printf("\t%10s%10.2f\n", size, price)
		}
	}
}
