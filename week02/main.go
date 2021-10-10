package main

import "fmt"

func main() {
	// snippet: main -- An array
	names := [4]string{}
	names[0] = "John"
	names[1] = "Paul"
	names[2] = "George"
	names[3] = "Ringo"
	// snippet: main

	fmt.Println(names)

	i := 0
	for i < len(names) {
		if i == 3 {
			break
		}
		fmt.Println(names[i])
		i++
	}
	fmt.Println()
	fmt.Println()
	//Slices
	nameSlice := []string{"Lou", "John", "Sterling", "Moe"}
	nameSlice = append(nameSlice, "Nico")
	for i, v := range nameSlice {
		fmt.Printf("%d. %s", i, v)
		fmt.Println()
	}
	fmt.Println()
	//maps
	beatles := map[string]string{}

	beatles["John"] = "guitar"
	beatles["Paul"] = "bass"
	beatles["George"] = "guitar"
	beatles["Ringo"] = "drums"

	fmt.Println(beatles)
	fmt.Println()
	for key, value := range beatles {
		fmt.Printf("%s plays %s\n", key, value)
	}

}
