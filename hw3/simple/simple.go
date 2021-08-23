package main

import (
	"fmt"
)

func main() {
	fmt.Println("Введите число N:")
	var n int
	_, err := fmt.Scan(&n)

	if err != nil {
		fmt.Println("Введите число")
		return
	}

	simples := getSimpleNumbers(n)
	for _, p := range simples {
		fmt.Println(p)
	}
}

func getSimpleNumbers(N int) (primes []int) {
	b := make(map[int]bool, N)

	for i := 2; i < N; i++ {
		if b[i] {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < N; k += i {
			b[k] = true
		}
	}

	return
}
