package main

import "fmt"

type person struct {
	name string
	age int
}

func main()  {

	fmt.Println(person{"Faith", 25})

	// You can name the fields when initializing a struct.
	fmt.Println(person{name:"Rhona", age:20})

	// Omitted fields will be zero-valued.
	fmt.Println(person{name: "Fred"})

	// An & prefix yields a pointer to the struct.
	fmt.Println(&person{name: "Ann", age: 40})

	// Access struct fields with a dot.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// You can also use dots with struct pointers - the pointers are automatically dereferenced.
	sp := &s
	sptr := *sp
	fmt.Println(sp.age)
	fmt.Println(sptr.age)

	// Structs are mutable.
	sp.age = 51
	fmt.Println(sp.age)

}


