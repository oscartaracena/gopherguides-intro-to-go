package main

import (
	"fmt"
	"testing"
)

//array test code
func TestArray(t *testing.T) {
	exp := [4]string{"John", "Paul", "George", "Ringo"}
	act := [4]string{}

	for i, v := range exp {
		act[i] = v
		fmt.Printf("%d - %s\n", i, v)
	}
	for i, v := range act {
		if act[i] != exp[i] {
			t.Errorf("act[%d] : %s is not equal to exp[%d] : %s ", i, v, i, exp[i])
		}
	}

}

//testing function for Slice

func TestSlice(t *testing.T) {

	exp := []string{"five", "six", "seven", "eight"}
	act := []string{"five", "six", "seven", "eight"}
	fmt.Println("Testing ....")

	for i, v := range exp {
		act[i] = v
		fmt.Printf("%d - %s\n", i, v)
	}
	for i, v := range act {
		if act[i] != exp[i] {
			t.Errorf("act[%d] : %s is not equal to exp[%d] : %s ", i, v, i, exp[i])
		}
	}
	fmt.Println("End Testing")
}

//testing function for map
func TestMap(t *testing.T) {
	exp := map[string]string{
		"Anthony K": "guitar",
		"Flea":      "bass",
		"John F":    "guitar",
		"Chad":      "drums",
	}

	act := map[string]string{
		"Anthony K": "guitar",
		"Flea":      "bass",
		"John F":    "guitar",
		"Chad":      "drums",
	}

	for key, value := range exp {
		act[key] = value
		fmt.Printf("%s - %s\n", key, value)
	}

	value, ok := exp["Flea"]
	if ok {
		fmt.Println("Key found value is: ", value)
	} else {
		fmt.Println("Key not found")
	}

}

/*
	key := "Anthony K"
	value, ok := exp[key]
	if !ok {
		return fmt.Errorf("Key not found: %q", key) // Key not found: "foo"
	}
	fmt.Printf("Found key %q: %q", key, value) // Found Key "Paul": "bass" */
