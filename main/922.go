package main

func sortArrayByParityII(nums []int) []int {
	i, j := 0, 1
	n := len(nums)
	for i < n && j < n {
		for i < n && nums[i]&1 == 0 {
			i += 2
		}
		for j < n && nums[j]&1 == 1 {
			j += 2
		}
		if j >= n || i >= n {
			return nums
		}
		nums[j], nums[i] = nums[i], nums[j]
		i += 2
		j += 2
	}
	return nums

}
