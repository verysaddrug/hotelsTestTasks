package main

import (
	"fmt"
	"strconv"
)

func main() {
	//test() // тест кейсы

	var num int
	fmt.Scan(&num)
	fmt.Println(computeDeclension(num))

}

func test() {
	fmt.Println(computeDeclension(1))
	fmt.Println(computeDeclension(2))
	fmt.Println(computeDeclension(5))
	fmt.Println(computeDeclension(21))
	fmt.Println(computeDeclension(25))
	fmt.Println(computeDeclension(41))
	fmt.Println(computeDeclension(1048))
}

func computeDeclension(number int) string {
	lastDigit := number % 10
	lastTwoDigits := number % 100

	// Определяем склонение в зависимости от числа
	switch {
	case lastTwoDigits >= 11 && lastTwoDigits <= 14:
		return strconv.Itoa(number) + " компьютеров"
	case lastDigit == 1:
		return strconv.Itoa(number) + " компьютер"
	case lastDigit >= 2 && lastDigit <= 4:
		return strconv.Itoa(number) + " компьютера"
	default:
		return strconv.Itoa(number) + " компьютеров"
	}
}
