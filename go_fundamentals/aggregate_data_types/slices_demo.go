package main

import (
	"fmt"
	"slices"
)

func slices_demo() {

	var s []string
	fmt.Println(s)
	s = []string{"Coffee", "Espresso", "Capuccino"}
	fmt.Println(s)
	fmt.Println(s[0])
	s[1] = "Chai"
	fmt.Println(s)

	s2 := s // reference
	s2[2] = "Chai Latte"
	fmt.Println(s, s2)

	s = append(s, "hot choco", "hot milk")
	fmt.Println(s)

	slices.Delete(s, 1, 2)
}
