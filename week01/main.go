package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello World")
	//Printing todo
	fmt.Println("Printing, TODO")
	a := "Say \"hello\" to Go!"
	fmt.Println(a)
	fmt.Printf("%s\n", a)

	b := "Hello, 世界"
	for i, c := range b {
		fmt.Printf("%d: %s\n", i, string(c))
	}

}
