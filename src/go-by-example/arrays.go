package main

import (
	"fmt"
	"reflect"
)

func main()  {
	var a [5] int
	fmt.Println("emp: ", a)
	fmt.Println(reflect.TypeOf(a).Kind())

	a[4] = 100
	a[0] = 20
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])

	fmt.Println("len: ", len(a))

	b := [5] int {1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	fmt.Println(reflect.TypeOf(b).Kind())

	var twoD [2][3] int
	for i := 0; i < 2; i ++ {
		for j := 0; j < 3; j ++ {
			twoD[i][j] = i + j

		}
	}
	fmt.Println("2d: ", twoD)
}