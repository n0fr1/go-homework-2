package bubblesort

import "fmt"

func SliceToSort() {

	sliceNum := []int{22, -1, 6, 9, 10, 0, 4, 1, -3, 8, 7, 13, 48}

	sortedSlice := Bubble(sliceNum)

	fmt.Printf("%s %d \n", "unsorted slice: ", sliceNum)
	fmt.Printf("%s %d", "sorted slice: ", sortedSlice)

}

//Bubble is slice-sorting function
func Bubble(sliceNum []int) []int {

	var sliceSortNum = make([]int, len(sliceNum))
	copy(sliceSortNum, sliceNum)

	for j := 1; j < len(sliceSortNum); j++ {
		for el := 0; el < len(sliceSortNum)-j; el++ {

			if sliceSortNum[el] > sliceSortNum[el+1] {
				sliceSortNum[el], sliceSortNum[el+1] = sliceSortNum[el+1], sliceSortNum[el]
			}
		}

	}

	return sliceSortNum
}
