package leetcode

import "fmt"

func main() {

	nums := []int{
		1,
		2,
		7,
		17,
		27,
	}

	target := 9

	tooSum(nums, target);
}

func tooSum(nums []int, target int) []int {
	var targetNums []int

	for value := range nums {
		fmt.Println(append(targetNums, target-value))

	}

	return []int{
		1,
		2,
	}

}
