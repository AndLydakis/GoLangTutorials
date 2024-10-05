package main

import "fmt"

func map_types_demo() {
	var m map[string][]string
	fmt.Println(m)

	m = map[string][]string{
		"coffee": {"Coffee", "espresso", "capuccino"},
		"tea":    {"Tea", "Chai", "Chai Latte"},
	}
	fmt.Println(m)

	fmt.Println(m["coffee"])

	m["other"] = []string{"Hot Chip"}
	fmt.Println(m)

	delete(m, "tea")
	fmt.Println(m)
	fmt.Println(m["tea"])

	v, ok := m["tea"]
	fmt.Println(v, ok)

	m2 := m
	fmt.Println(m2)
	delete(m, "coffee")
	fmt.Println(m, m2)
}
