package main

import (
	"fmt"
)

// Функция для вывода таблицы умножения
func printMultiplicationTable(n int) {
	// Используем вложенные циклы для генерации строк и столбцов таблицы умножения
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}
}

func main() {
	var n int
	fmt.Println("Введите число для генерации таблицы умножения:")
	fmt.Scanln(&n)

	printMultiplicationTable(n)
}
