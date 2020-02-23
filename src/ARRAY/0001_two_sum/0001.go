package _001_two_sum

/**
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

func twoSum(nums []int, target int) []int {
	var result = make([]int, 0)
	var m = make(map[int]int)
	for i, val := range nums {
		idx, ok := m[target - val]
		if ok {
			result = []int{i, idx}
			break
		} else {
			m[val] = i
		}
	}
	return result
}
