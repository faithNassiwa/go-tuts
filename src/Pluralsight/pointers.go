package main

import "fmt"

func main()  {
	name := "Faith"
	course := "Go Fundamentals"
	module := 4.5
	ptr_name := &name
	ptr_course := &course
	ptr_module := &module

	fmt.Println("The memory address of *name* variable is ", ptr_name, "and the value of *name* is", *ptr_name)
	fmt.Println("The memory address of *course* variable is ", ptr_course, "and the value of *course* is",
		*ptr_course)
	fmt.Println("The memory address of *module* variable is ", ptr_module, "and the value of *module* is",
		*ptr_module)

}
