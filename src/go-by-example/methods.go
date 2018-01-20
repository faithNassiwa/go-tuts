package main

import "fmt"

type rect struct {
	width, height int
}

// This area method has a pointer receiver.
func (r *rect) area() int  {
	return r.width * r.height
}

// This perim method has a value receiver.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main()  {
	r := rect{10, 5}
	fmt.Println("Area: ", r.area())
	fmt.Println("Perimeter: ", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())

}