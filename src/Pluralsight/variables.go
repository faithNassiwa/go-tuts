package main

import (
	"fmt"
	"reflect"
	)

var (
	//name, course string
	//module float64

	name, course, module = "Faith", "Go Fundamentals", 4.3
)
func main()  {
	fmt.Println("Name is set to", name)
	fmt.Println("Module is set to", module)
	fmt.Println("Name is ", name,  "and is of type", reflect.TypeOf(name))
	fmt.Println("Course is ", course,  "and is of type", reflect.TypeOf(course))
	fmt.Println("Module is ", module,  "and is of type", reflect.TypeOf(module))

	a := 10.0000  // shorthand(:= ) construct only works in functions and when initializing and declaring variables
	b := 3
	fmt.Println("\nA is of type", reflect.TypeOf(a), "and B is of type ", reflect.TypeOf(b))

	// c := a + b // invalid operation: a + b (mismatched types float64 and int)

	c := int(a) + b // type casting a to type int for it to be added to b

	//d := 0  variables declared at function level must be used

	fmt.Println("\nC has a value ", c, "and is of type ", reflect.TypeOf(c))


}
