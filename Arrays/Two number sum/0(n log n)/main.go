//0(nlog(n)) time  0(1) space
//This is the best solution to solve the problem of two number sum
package main

import "sort"

func TwoNumberSum(array []int, target int) []int {

	sort.Ints(array)
	left, right := 0, len(array)-1
	for left < right {
		sum := array[left] + array[right]
		if sum == target {
			return []int{array[left], array[right]}
		} else if sum < target {
			left += 1
		} else {
			right -= 1

		}
	}

	return []int{}
}
