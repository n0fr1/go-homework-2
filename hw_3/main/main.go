package main

import (
	"fmt"
	"go-homework-level2/hw_3/createFiles"
	"go-homework-level2/hw_3/slicePanic"
)

func main() {

	//1. panic - out of range
	slicePanic.RunPanic(10)
	fmt.Printf("%s\n", "Program continue...")

	//2. own error
	err := slicePanic.New("some error")
	fmt.Printf("%v\n", err)

	//3. сreate 5 files
	createFiles.Newfiles()
}
