package createFiles

import (
	"fmt"
	"os"
	"strconv"
)

func Newfiles() {

	for i := 0; i < 1000001; i++ {

		func() {

			iString := strconv.Itoa(i)
			file, err := os.Create("new" + iString + ".go")

			if err != nil {
				fmt.Println("file wasn't created!")
			}

			defer file.Close()
		}()
	}

}
