package main

import (
	"fmt"

	"github.com/n0fr1/go-homework-2/hw_3/bubblesort"
	"github.com/n0fr1/go-homework-2/hw_3/slicePanic"
)

func main() {

	//1. panic - out of range
	slicePanic.RunPanic(10)
	fmt.Printf("%s\n", "Program continue...")

	//2. own error
	err := slicePanic.New("some error")
	fmt.Printf("%v\n", err)

	//3. bubble sorting for slice
	bubblesort.SliceToSort()
}
