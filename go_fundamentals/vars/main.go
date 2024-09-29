package main

import "fmt"

func main() {
	a := "This is an interpreted string"
	// b := 'This is a raw string'
	fmt.Println(a)

	b := -1 // int
	fmt.Println(b)

	c := true //bool
	fmt.Println(c, !c)

	var myName string = "explicit test string declaration"
	fmt.Println(myName)

	var myName2 = "infered type"
	myName3 := "short declaration"

	fmt.Println(myName2)
	fmt.Println(myName3)

	var i int = 32
	var f, ff float32
	// f = i

	ff = float32(i)

	fmt.Println(f, ff)

	// Arithemtic OPS
	op1, op2 := 10, 4
	op3 := op1 + op2
	op3 = op1 - op2
	op3 = op1 * op2
	op3 = op1 / op2
	op3 = op1 / 3
	op3 = op1 % 3
	d := 7.0 / 2.0
	fmt.Println(op3, d)

	fmt.Println(op1 == op2)
	fmt.Println(op1 != op2)
	fmt.Println(op1 > op2)

	fmt.Println(op1 / op2)
	fmt.Println(float32(op1) / float32(op2))

	// Constants
	fmt.Println("---- constants")
	const c1 = 42
	const c2 string = "const string"

	fmt.Println(c1, c2)

	const c3 = c1
	fmt.Println(c3)

	fmt.Println("---- group constants")
	const (
		c4 = true
		c5 = 3.14
	)
	fmt.Println(c4, c5)

	fmt.Println("----- implicit constants")
	const (
		c6 = "foo"
		c7 //"foo"
	)
	fmt.Println(c6, c7)

	fmt.Println("------ iota")
	const i1 = iota // 0
	fmt.Println(i1)

	fmt.Println("---- group iota")
	const (
		i2 = iota // 0
		i3        //  1
		i4        // 2
		i5        // 3
	)
	fmt.Println(i2, i3, i4, i5)

	fmt.Println("---- pointers")
	p1 := 24
	p2 := &p1
	fmt.Println(p1, p2, *p2)
	p1 = 25
	fmt.Println(p1, p2, *p2)

	p3 := "f00"
	p4 := &p3
	fmt.Println(p3, p4)
	*p4 = "bar"
	fmt.Println(p3, p4)

	fmt.Println("----- pointer demo")
	p5 := "Hello, world"
	p6 := &p5
	fmt.Println(p5, p6)
	fmt.Println(*p6)
	*p6 = "Helloooo,world"
	fmt.Println(*p6, p6)

	p6 = new(string) //still needs to point to a string
	fmt.Println(p6, "-->", *p6)
}
