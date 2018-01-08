package main

import (
	"fmt"
	"runtime"
)

// Prints hello from os you are running
func main()  {
	fmt.Println("Hello from", runtime.GOOS)
	
}
