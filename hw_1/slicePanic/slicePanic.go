package slicePanic

import (
	"fmt"
	"log"
	"time"
)

type ErrorTexTime struct {
	text string
	time string
}

func RunPanic(index int) {

	defer func() { //catch error - out of range

		if err := recover(); err != nil {
			var inf ErrorTexTime
			inf.time = time.Now().Format(time.RFC3339)
			inf.text = "slice - out of range"
			log.Printf("panic found: %s, %v", inf.text, inf.time)
		}

	}()

	slice := []int{1}
	fmt.Println(slice[index])

}

func (e *ErrorTexTime) Error() string {
	return fmt.Sprintf("Error %v at %v", e.text, e.time)
}

func New(text string) error {
	return &ErrorTexTime{
		text: text,
		time: time.Now().Format(time.RFC3339),
	}
}
