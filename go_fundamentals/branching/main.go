package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

func if_demo() {
	i := 5
	if i < 5 {
		fmt.Println("less than 5")
	} else if i < 10 {
		fmt.Println(" < 10")
	} else {
		fmt.Println("at least 10")
	}

}

func switch_demo() {
	i := 5
	switch i {
	case 1:
		fmt.Println("first cast")
	case 2 + 3, 2*i + 3:
		fmt.Println("Second case")
	default:
		fmt.Println("default")
	}

}

func deferred_demo() {
	db, _ := sql.Open("driverName", "connection string")
	defer db.Close()

	rows, _ := db.Query("some sql query here")
	defer rows.Close()
}

func func1(dividend, divisor int) int {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg)
		}
	}()
	return dividend / divisor

}
func panic_recover_demo() {
	fmt.Println("-----")
	dividend, divisor := 10, 5
	fmt.Printf("%v divided by %v is %v\n", dividend, divisor, func1(dividend, divisor))
	dividend, divisor = 10, 0
	fmt.Printf("%v divided by %v is %v\n", dividend, divisor, func1(dividend, divisor))
	fmt.Println("------")
}

type MenuItem struct {
	name   string
	prices map[string]float64
}

func main() {
	// if_demo()
	// switch_demo()

	panic_recover_demo()
	menu := []MenuItem{
		{name: "Coffee", prices: map[string]float64{"small": 1.65, "medium": 2.5, "large": 3.0}},
		{name: "Espresso", prices: map[string]float64{"small": 1.65, "medium": 2.5, "large": 3.0}},
	}
loop:
	for {
		fmt.Println("Select an option")
		fmt.Println("1) Print Menu")
		fmt.Println("2) Add Item")
		fmt.Println("q) Quit")
		in := bufio.NewReader(os.Stdin)
		choice, _ := in.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			for _, v := range menu {
				fmt.Println(v.name)
				fmt.Println(strings.Repeat("-", 10))
				for size, price := range v.prices {
					fmt.Printf("\t%10s%10.2f\n", size, price)
				}
			}
		case "2":
			fmt.Println("Please enter new item")
			name, _ := in.ReadString('\n')
			menu = append(menu, MenuItem{name: name, prices: make(map[string]float64)})
		case "q":
			fmt.Println("Quit")
			break loop
		default:
			break loop
		}
	}

}
