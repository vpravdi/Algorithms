package main

func TwoNumberSum(array []int, target int) []int {

	m := map[int]bool{}
	for _, v := range array {
		form := target - v
		if _, ok := m[form]; ok {
			return []int{form, v}
		}
		m[v] = true
	}
	return []int{}
}
