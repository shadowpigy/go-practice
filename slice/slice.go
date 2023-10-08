package main

import "fmt"

func Delete[Type any](s []Type, n int) []Type {
	if n < 0 || n >= len(s) {
		panic("slice index out of bounds")
	}

	var numElement = len(s) - 1

	if len(s) >= 16 && len(s)-1 < cap(s)>>2 {

		var res = make([]Type, numElement, cap(s)>>1+1)
		return deleteHelper(s, n, res)
	}

	var res = make([]Type, numElement, cap(s))
	return deleteHelper(s, n, res)
}

func deleteHelper[Type any](s []Type, n int, res []Type) []Type {
	for i := 0; i < n; i++ {
		res[i] = s[i]
	}
	for i := n; i < len(s)-1; i++ {
		res[i] = s[i+1]
	}
	return res
}

func main() {
	var arr = []int{1, 2, 3, 4, 5}
	fmt.Printf("%+v\n", Delete(arr, 0))
}
