package main

import "fmt"

func Delete[Type any](s []Type, n int) ([]Type, bool) {
	if n < 0 || n >= len(s) {
		return []Type{}, false
	}

	var numElement = len(s) - 1

	if len(s) >= 16 && len(s)-1 < cap(s)>>2 {
		var res = make([]Type, numElement, cap(s)>>1)
		return deleteHelper(s, n, res), true
	}

	var res = make([]Type, numElement, cap(s))
	return deleteHelper(s, n, res), true
}

func deleteHelper[Type any](s []Type, n int, res []Type) []Type {
	copy(res[:n], s[:n])
	copy(res[n:], s[n+1:])
	return res
}

func main() {
	var arr = []int{1, 2, 3, 4, 5}
	res, _ := Delete(arr, 0)
	fmt.Printf("%+v\n", res)
}
