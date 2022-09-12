package main

func specialArray(nums []int) int {
	count := make([]int, 10001)
	for _, num := range nums {
		count[num]++
	}
	sum := 0
	for i := 1000; i >= 0; i-- {
		sum += count[i]
		if i == sum {
			return sum
		} else if i < sum {
			return -1
		}
	}
	return -1
}
