package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// Преобразуем римские числа в арабские
func romanToArabic(roman string) (int, error) {
	romanNumerals := map[string]int{
		"M": 1000, "CM": 900, "D": 500, "CD": 400,
		"C": 100, "XC": 90, "L": 50, "XL": 40,
		"X": 10, "IX": 9, "V": 5, "IV": 4, "I": 1,
	}
	result := 0
	i := 0

	for i < len(roman) {
		if i+1 < len(roman) && romanNumerals[roman[i:i+2]] > 0 {
			result += romanNumerals[roman[i:i+2]]
			i += 2
		} else if romanNumerals[roman[i:i+1]] > 0 {
			result += romanNumerals[roman[i:i+1]]
			i++
		} else {
			return 0, fmt.Errorf("некорректное римское число")
		}
	}

	if !validateRoman(roman) {
		return 0, fmt.Errorf("некорректное римское число")
	}

	if result < 1 || result > 10 {
		return 0, fmt.Errorf("число вне допустимого диапазона 1-10")
	}
	return result, nil
}

func validateRoman(roman string) bool {
	pattern := `^M{0,4}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$`
	match, _ := regexp.MatchString(pattern, roman)
	return match
}

func arabicToRoman(arabic int) string {
	arabicNumerals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanNumerals := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	result := ""

	for i, arabicValue := range arabicNumerals {
		for arabic >= arabicValue {
			arabic -= arabicValue
			result += romanNumerals[i]
		}
	}
	return result
}

func main() {
	var input1, input2, operation string

	fmt.Print("Введите выражение (например, 2 + 2 или II + III): ")
	fmt.Scanf("%s %s %s", &input1, &operation, &input2)

	isRoman := func(s string) bool {
		for _, char := range s {
			if _, ok := map[rune]int{'I': 1, 'V': 5, 'X': 10}[char]; !ok {
				return false
			}
		}
		return true
	}

	if isRoman(input1) && isRoman(input2) {
		a, err1 := romanToArabic(input1)
		b, err2 := romanToArabic(input2)
		if err1 != nil || err2 != nil {
			fmt.Printf("Ошибка: %v\n", err1)
			return
		}
		processOperation(a, b, operation, true)
	} else {
		a, err1 := strconv.Atoi(input1)
		b, err2 := strconv.Atoi(input2)
		if err1 != nil || err2 != nil {
			fmt.Printf("Ошибка: %v\n", err1)
			return
		}
		if a < 1 || a > 10 || b < 1 || b > 10 {
			fmt.Println("Ошибка: Числа должны быть в диапазоне от 1 до 10.")
			return
		}
		processOperation(a, b, operation, false)
	}
}

func processOperation(a, b int, operation string, isRoman bool) {
	var result int
	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("На ноль делить нельзя!")
			return
		}
		result = a / b
	default:
		fmt.Println("Неверная операция. Доступные операции: +, -, *, /.")
		return
	}

	if isRoman {
		if result < 1 {
			fmt.Println("Ошибка: Результат римской арифметики не может быть меньше 1.")
			return
		}
		fmt.Println(arabicToRoman(result))
	} else {
		fmt.Println(result)
	}
}
