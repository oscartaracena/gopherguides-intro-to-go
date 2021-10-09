package main

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	exp := [4]string{"John", "Paul", "George", "Ringo"}
	act := [4]string{}

	for i, v := range exp {
		act[i] = v
	}
}
