package main

import (
	"fmt"
	"go-homework-level2/hw_1/createFiles"
	"go-homework-level2/hw_1/slicePanic"
)

func main() {

	//1. panic - out of range
	slicePanic.RunPanic(10)
	fmt.Println("Program continue...")

	//2. own error
	err := slicePanic.New("some error")
	fmt.Println(err)

	//3. сreate 5 files
	createFiles.Newfiles()
}
