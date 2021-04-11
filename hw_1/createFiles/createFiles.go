package createFiles

import (
	"fmt"
	"os"
	"strconv"
)

func Newfiles() {

	for i := 0; i < 1000000; i++ {

		func() {

			iString := strconv.Itoa(i)
			file, err := os.Create(iString + ".go")

			if err != nil {
				fmt.Println("file wasn't created!")
				return
			}

			defer file.Close()
		}()
	}

}
