/* Written by Keith Drew, 2018
* This is a solution to LeetCode problem #1: "Two Sum," written in Go
*
* The answer could be better. Some improvements:
* - Handle the case of no solution better.
* - Handle the case of multiple solutions.
* - Optimize the code.
* - Read array and target values from STDIN to more quickly test many inputs.
* - ??
*/
package main

import "fmt"

func main() {
     	nums := []int{2, 7, 11, 15}
	var target int = 13

	fmt.Println(nums)
	fmt.Println(target)

	var solution []int = twoSum(nums, target)

	fmt.Println("Solution: ", solution)

}

func twoSum(nums []int, target int) []int {
	indexes := make([]int, 2)

	length := len(nums)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				indexes[0] = i
				indexes[1] = j
			}
		}
	}
	return indexes
}
