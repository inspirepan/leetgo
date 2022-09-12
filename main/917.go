package main

import (
	"unicode"
)

func reverseOnlyLetters(s string) string {
	l, r := 0, len(s)
	arr := []rune(s)
	for l < r {
		for l < r && !unicode.IsLetter(arr[l]) {
			l++
		}
		for l < r && !unicode.IsLetter(arr[r]) {
			r--
		}
		if l >= r {
			break
		}
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
	return string(arr)
}
