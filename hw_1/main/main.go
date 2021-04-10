package main

import (
	"fmt"
	"go-homework-level2/createFiles"
	"go-homework-level2/slicePanic"
)

func main() {

	//1. panic - out of range
	slicePanic.RunPanic(10)
	fmt.Println("Program continue...")

	slicePanic.RunPanic(0)
	fmt.Println("Program continue...")

	//2. сreate 1 000 000 files
	createFiles.Newfiles()
}
