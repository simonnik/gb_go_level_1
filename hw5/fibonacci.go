package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Введите число:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, err := strconv.ParseInt(scanner.Text(), 10, 64)

	if err != nil {
		fmt.Println(err)

		return
	}
	fmt.Println(fibonacciWithMap(input, map[int64]int64{1: 1, 2: 1}))
}

func fibonacciWithMap(n int64, cache map[int64]int64) int64 {
	if val, found := cache[n]; found {
		return val
	}

	cache[n] = fibonacciWithMap(n-1, cache) + fibonacciWithMap(n-2, cache)

	return cache[n]
}
