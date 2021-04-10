package main

import (
	"fmt"
	"go-homework-level2/createFiles"
	"go-homework-level2/slicePanic"
)

func main() {

	//out of range
	slicePanic.RunPanic(10)
	fmt.Println("Program continue...")

	slicePanic.RunPanic(0)
	fmt.Println("Program continue...")

	//сreate 1 000 000 files
	createFiles.Newfiles()
}
