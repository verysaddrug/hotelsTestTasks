package main

import (
	"fmt"
)

// Функция для проверки, является ли число простым
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// Функция для нахождения всех простых чисел в диапазоне [min, max]
func primesInRange(min, max int) []int {
	var primes []int
	for i := min; i <= max; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	var min, max int
	fmt.Println("Введите минимальное число диапазона:")
	fmt.Scanln(&min)
	fmt.Println("Введите максимальное число диапазона:")
	fmt.Scanln(&max)

	result := primesInRange(min, max)
	fmt.Println("Простые числа в диапазоне:", result)
}

