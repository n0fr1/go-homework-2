package createFiles

import (
	"fmt"
	"os"
	"strconv"
)

func Newfiles() {
	func() {
		for i := 0; i < 5; i++ {
			iString := strconv.Itoa(i)
			file, err := os.Create(iString + ".go")

			if err != nil {
				fmt.Println("file " + file.Name() + " wasn't created!")
				return
			}

			defer file.Close()
		}
	}()

}
