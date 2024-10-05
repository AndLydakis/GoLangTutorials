package main

import "fmt"

func arrays() {
	var arr [3]int
	fmt.Println(arr)
	arr = [3]int{1, 2, 3}
	fmt.Println(arr[1])
	fmt.Println(arr)
	fmt.Println(len(arr))

	sarr := [3]string{"f", "b", "bz"}
	sarr2 := sarr //copy
	fmt.Println(sarr)
	sarr[0] = "g"
	fmt.Println(sarr2)

	fmt.Println(sarr == sarr2)

	fmt.Println("--------------")
	var arr_ [3]string
	fmt.Println(arr_)
	arr_ = [3]string{"Coffee", "Espresso", "Capuccino"}
	fmt.Println(arr_)
	fmt.Println(arr_[0])
	arr_[1] = "Chai"
	fmt.Println(arr_)

	arr2 := arr_
	arr2[2] = "Chai Latte"
	fmt.Println(arr_, arr2)
}
func main() {
	// arrays()
	// slices_demo()
	// map_types_demo()
	structs_demo()

}
