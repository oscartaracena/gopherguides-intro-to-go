package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello World")
	//Printing todo
	fmt.Println("Printing, TODO")
	
	//printing String ("Go")
	str := "Go"
	fmt.Println("Printing", str)
	
	//printing int (42) 
	numVal := 42
	fmt.Println("Printing", numVal)
	
	//printing bool (true)
	
	fmt.Println("Printing")	

	//ignore just testing that is all 
	a := "Say \"hello\" to Go!"
	fmt.Println(a)
	fmt.Printf("%s\n", a)
	b := "Hello, 世界"
	for i, c := range b {
		fmt.Printf("%d: %s\n", i, string(c))
	}

}
