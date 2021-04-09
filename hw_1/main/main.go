package main

import (
	"fmt"
	"go-homework-level2/slicePanic"
)

func main() {
	slicePanic.RunPanic(10)
	fmt.Println("Program continue...")

	slicePanic.RunPanic(0)
	fmt.Println("Program continue...")
}
