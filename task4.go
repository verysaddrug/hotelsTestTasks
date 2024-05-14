package main

import (
    "fmt"
)

// Функция для вывода таблицы умножения
func printMultiplicationTable(n int) {
    // Выводим заголовок таблицы
    fmt.Printf("    ")
    for i := 1; i <= n; i++ {
        fmt.Printf("%4d", i)
    }
    fmt.Println()

    // Выводим разделительную линию
    fmt.Printf("    ")
    for i := 1; i <= n; i++ {
        fmt.Printf("----")
    }
    fmt.Println()

    // Выводим строки таблицы умножения
    for i := 1; i <= n; i++ {
        fmt.Printf("%4d|", i) // Выводим номер строки
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
