package main

import "fmt"

func main() {

	nums := []int{
		1,
		2,
		7,
		17,
		27,
	}

	target := 8

	fmt.Println(tooSum(nums, target))
}

func tooSum(nums []int, target int) []int {
	//var targetNums []int
	numLen := len(nums)

	for index, value := range nums {

		for j := index + 1; j < numLen; j++ {

			if value+nums[j] == target {
				return []int{
					index,
					j,
				};
			}

		}

	}

	return nil

}
