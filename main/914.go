package main

func hasGroupsSizeX(deck []int) bool {
	countMap := make(map[int]int)
	for _, i := range deck {
		countMap[i]++
	}
	var x = -1
	for _, val := range countMap {
		if x == 1 {
			return false
		} else {
			x = gcd(x, val)
		}
	}
	return x > 1
}

func gcd(x int, y int) int {
	if x == -1 {
		return y
	}
	for x > 0 && y > 0 {
		if x > y {
			x = x % y
		} else {
			y = y % x
		}
	}
	if x == 0 {
		return y
	}
	return x
}
