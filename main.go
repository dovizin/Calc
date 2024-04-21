package main

import "fmt"

func main() {
	// Объявляем переменные
	dividend := 9
	divisor := 4

	// Выполняем деление и выводим остаток от деления
	остаток := dividend % divisor
	fmt.Println("Остаток от деления 9 на 4:", остаток)
}
