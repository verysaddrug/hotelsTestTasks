package main

import (
	"fmt"
)

// Функция для нахождения наибольшего общего делителя (НОД) двух чисел
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Функция для нахождения всех делителей числа
func findDivisors(n int) []int {
	var divisors []int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}

// Функция для нахождения общих делителей массива чисел
func commonDivisors(numbers []int) []int {
	if len(numbers) == 0 {
		return []int{}
	}

	// Находим НОД всех чисел в массиве
	commonGCD := numbers[0]
	for _, num := range numbers[1:] {
		commonGCD = gcd(commonGCD, num)
		if commonGCD == 1 {
			return []int{1} // 1 - единственный общий делитель, если больше никаких нет
		}
	}

	// Находим все делители НОД
	divisors := findDivisors(commonGCD)
	return divisors
}

func main() {

	//test() // testcase
	
	var num int
	fmt.Println("Введите количество элементов массива:")
	fmt.Scanln(&num)

	nums := make([]int, num)

	fmt.Println("Введите элементы массива:")
	for i := 0; i < num; i++ {
		fmt.Scan(&nums[i])
	}

	result := commonDivisors(nums)
	fmt.Println(result)
}

func test() {
	nums := []int{42, 12, 18}
	result := commonDivisors(nums)
	fmt.Println(result)
}
