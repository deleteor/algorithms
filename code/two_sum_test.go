package code

import "testing"

/*
两数之和（简单）
http://leetcode-cn.com/problems/two-sum
*/
func TestTwoSum(t *testing.T) {
	var (
		nums   = []int{2, 7, 11, 15}
		target = 9
	)

	t.Log(twoSum(nums, target))
	t.Log(twoSum1(nums, target))
}

// 双重循环
func twoSum(nums []int, target int) []int {
	var (
		lens = len(nums)
		out  = make([]int, 2)
	)

	if lens < 2 {
		return out
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				out[0] = i
				out[1] = j
				return out
			}
		}
	}
	return out
}

// 一次循环  利用map唯一key
// key为 target-nums[i]  value记录i 下表
func twoSum1(nums []int, target int) []int {
	result := []int{}
	m := make(map[int]int)
	for i, k := range nums {
		if value, exist := m[target-k]; exist {
			result = append(result, value)
			result = append(result, i)
		}
		m[k] = i
	}
	return result
}
