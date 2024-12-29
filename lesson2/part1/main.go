package main

import "fmt"

// slice

// || - OR

// && - AND

// ! - NOT

func popFromSlice(arr []int, index int) ([]int, error) {
	if index < 0 || index >= len(arr) {
		return arr, fmt.Errorf("invalid index")
	}

	return append(arr[:index], arr[index+1:]...), nil // None
}

func main() {
	a := []int{123123, 112, 2332, 513}
	a, err := popFromSlice(a, 2)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(a)
	}
}
