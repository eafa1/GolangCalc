package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RomanToArabic(num string) (int, error) { // конвертация римского числа в арабское
	romanNumerals := map[string]int{"I": 1, "V": 5, "X": 10}
	arabic := 0
	lastValue := 0

	for i := len(num) - 1; i >= 0; i-- {
		value, exists := romanNumerals[string(num[i])]
		if !exists {
			return 0, fmt.Errorf("неверное римское число")
		}
		if value < lastValue {
			arabic -= value
		} else {
			arabic += value
		}
		lastValue = value
	}
	return arabic, nil
}

func ArabicToRoman(num int) (string, error) { // конвертация арабского числа в римское
	if num < 1 {
		return "", fmt.Errorf("в римской системе нет нуля или отрицательных чисел")
	}

	values := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var result strings.Builder
	for i, value := range values {
		for num >= value {
			num -= value
			result.WriteString(symbols[i])
		}
	}
	return result.String(), nil
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Input:")

		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		chars := strings.Fields(text)

		if len(chars) != 3 {
			fmt.Println("Ошибка: неправильный формат ввода, калькулятор умеет выполнять операции сложения, вычитания, умножения и деления с двумя числами: a + b, a - b, a * b, a / b")
			break
		}

		num1, err1 := strconv.Atoi(chars[0])
		num2, err2 := strconv.Atoi(chars[2])

		if (err1 != nil) != (err2 != nil) {
			fmt.Println("Ошибка: калькулятор умеет работать только с целыми арабскими или римскими цифрами одновременно")
			break
		}

		if (err1 == nil) && (err2 == nil) {

		}

		isRoman := err1 != nil && err2 != nil
		var romanResult string
		var err error

		if isRoman {
			num1, err = RomanToArabic(chars[0])
			if err != nil {
				fmt.Println(err)
				break
			}
			num2, err = RomanToArabic(chars[2])
			if err != nil {
				fmt.Println("Ошибка: ", err)
				break
			}
		}

		if (num1 > 10) || (num2 > 10) {
			fmt.Println("Ошибка: калькулятор умеет работать только с целыми числами не больше 10")
			break
		}

		var result int

		switch chars[1] {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			if num2 == 0 {
				fmt.Println("Ошибка: деление на ноль")
				os.Exit(1)
			}
			result = num1 / num2
		default:
			fmt.Println("Ошибка: неверная арифметических операция")
			os.Exit(2)
		}

		if isRoman {
			romanResult, err = ArabicToRoman(result)
			if err != nil {
				fmt.Println("Ошибка: ", err)
				break
			}
			fmt.Printf("\nOutput:\n%v\n\n", romanResult)
		} else {
			fmt.Printf("\nOutput:\n%v\n\n", result)
		}
	}

}
