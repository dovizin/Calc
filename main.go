package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	romain_num := map[string]int{
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
	arabic_num := map[string]int{
		"0":  0,
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
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') // Считываем строку до символа новой строки
	checkvalue(input)
	num := strings.Fields(input)

	num1 := num[0]
	operator := num[1]
	num2 := num[2]

	issupset(num2, arabic_num, romain_num)
	issupset(num1, arabic_num, romain_num)
	f := valide(num1, num2, arabic_num, romain_num, operator)
	fmt.Println(f)
}

func checkvalue(num string) {
	operators := []string{"/", "*", "-", "+"}
	var countOperators int
	for _, opr := range operators {
		countOperators += strings.Count(num, opr)
	}
	if countOperators == 0 {
		fmt.Println(countOperators)
		panic("Выдача паники, так как строка не является математической операцией.")

	} else if countOperators > 1 {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}
}

var result int

func calculate(num1 int, num2 int, operand string, output bool) interface{} {
	switch operand {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			panic("нельзя делить на 0")
		}
		result = num1 / num2
	default:
		panic("неверный оператор")

	}
	if !output {
		return arabTorom(result)

	}
	return result
}
func arabTorom(value int) string {
	romain_num := map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
	}
	return romain_num[value]
}

func isExist(value string, list map[string]int) bool {
	_, exist := list[value]
	return exist
}
func romToArab(numstr string) int {
	var romain_num = map[string]int{
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
	num1 := romain_num[numstr]
	return num1
}
func issupset(num string, list map[string]int, list2 map[string]int) {
	if !isExist(num, list) && !isExist(num, list2) {
		panic("не подходящие числа")
	}
}
func valide(num1str string, num2str string, arabic_num map[string]int, romain_num map[string]int, operator string) interface{} {

	if isExist(num1str, arabic_num) && isExist(num2str, arabic_num) {
		num1, _ := strconv.Atoi(num1str)
		num2, _ := strconv.Atoi(num2str)
		result := calculate(num1, num2, operator, true)
		return result

	} else if isExist(num1str, romain_num) && isExist(num2str, romain_num) {
		num1 := romToArab(num1str)
		num2 := romToArab(num2str)
		if operator == "-" && num1 < num2 {
			panic("Выдача паники, так как в римской системе нет отрицательных чисел.")
		}
		result := calculate(num1, num2, operator, false)
		return result

	} else if isExist(num1str, arabic_num) && isExist(num2str, romain_num) || isExist(num1str, romain_num) && isExist(num2str, arabic_num) {
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
	}

	return result
}
