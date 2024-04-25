package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var num1, num2 string
	var input string

	romanNum := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	arabicNum := map[string]int{
		"1":  1,
		"2":  2,
		"3":  3,
		"4":  4,
		"5":  5,
		"6":  6,
		"7":  7,
		"8":  8,
		"9":  9,
		"10": 10,
	}

	fmt.Println("Введите выражение:")
	fmt.Scanln(&input)
	num := strings.Split(strings.TrimSpace(input), " ")

	if len(num) != 3 {
		fmt.Println("Некорректный формат ввода")
		return
	}

	num1 = num[0]
	operand := num[1]
	num2 = num[2]

	result, err := valide(num1, num2, arabicNum, romanNum, operand)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Результат:", result)
}

func calculate(num1, num2 int, operand string) (int, error) {
	switch operand {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("нельзя делить на 0")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("недопустимый операнд")
	}
}

func isExist(value string, list map[string]int) bool {
	_, ok := list[value]
	return ok
}

func valide(num1str, num2str string, arabicNum, romanNum map[string]int, operand string) (int, error) {
	if isExist(num1str, arabicNum) && isExist(num2str, arabicNum) {
		num1, _ := strconv.Atoi(num1str)
		num2, _ := strconv.Atoi(num2str)
		return calculate(num1, num2, operand)
	} else if isExist(num1str, romanNum) && isExist(num2str, romanNum) {
		num1 := romanNum[num1str]
		num2 := romanNum[num2str]
		return calculate(num1, num2, operand)
	}

	return 0, fmt.Errorf("не удалось выполнить операцию: числа или системы счисления не поддерживаются")
}
