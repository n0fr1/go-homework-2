package slicePanic

import (
	"fmt"
	"time"
)

type ErrorTexTime struct {
	text string
	time string
}

func RunPanic(index int) {

	defer func() { //catch error - out of range

		if err := recover(); err != nil {
			getErr := New("slice - out of range")
			fmt.Printf("%v\n", getErr)
		}

	}()

	slice := []int{1}
	fmt.Println(slice[index])

}

func (e *ErrorTexTime) Error() string {
	return fmt.Sprintf("Error: %v at %v", e.text, e.time)
}

func New(text string) error {
	return &ErrorTexTime{
		text: text,
		time: time.Now().Format(time.RFC3339),
	}
}
