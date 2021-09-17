package main

import "fmt"

func main() {
	res := insertion([]int{5, 3, 6, 8, 1, 2})
	fmt.Println(res)
}

func insertion(arr []int) []int {
	for i := 1; i <= len(arr)-1; i++ {
		x := arr[i]
		for j := i; j > 0 && arr[j-1] > x; j-- {
			arr[j], arr[j-1] = arr[j-1], x
		}
	}
	return arr
}
