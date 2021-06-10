//If any two numbers in the array sum up to target sum, the function would return them in an array

package main

func TwoNumberSum(array []int, target int) []int {

	for i := 0; i < len(array)-1; i++ {
		firstNum := array[i]
		for j := i + 1; j < len(array); j++ {
			secondNum := array[j]
			if firstNum+secondNum == target {
				return []int{firstNum, secondNum}
			}
		}
	}
	return []int{}
}
