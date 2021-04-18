package createFiles

import (
	"fmt"
	"os"
	"strconv"
)

func Newfiles() {

	for i := 0; i < 5; i++ {
		func() {

			iString := strconv.Itoa(i)
			file, err := os.Create(iString + ".go")

			if err != nil {
				fmt.Printf("file %s %s", file.Name(), " wasn't created!")
				return
			}

			defer file.Close()
		}()
	}

}
