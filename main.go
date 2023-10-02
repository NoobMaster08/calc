package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var romanNumerals = []struct {
	Value  int
	Symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func isRoman(s string) bool {
	return regexp.MustCompile(`^[IVXLCDM]+$`).MatchString(s)
}

func isArabic(s string) bool {
	return regexp.MustCompile(`^[0-9]+$`).MatchString(s)
}

func romanToArabic(roman string) (int, error) {
	total := 0
	for _, numeral := range romanNumerals {
		for strings.HasPrefix(roman, numeral.Symbol) {
			total += numeral.Value
			roman = roman[len(numeral.Symbol):]
		}
	}

	if total == 0 {
		return 0, fmt.Errorf("некорректное римское число")
	}
	return total, nil
}

func arabicToRoman(arabic int) string {
	if arabic <= 0 {
		return "Недопустимо для римских цифр"
	}
	var result strings.Builder
	for _, numeral := range romanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return result.String()
}

func calculator(input string) (string, error) {
	var a, b, result int
	var operation string
	var useRoman bool

	re := regexp.MustCompile("(\\d+|[IVXLCDM]+)\\s*([\\+\\-\\*/])\\s*(\\d+|[IVXLCDM]+)")
	matches := re.FindStringSubmatch(input)

	if len(matches) != 4 {
		return "", fmt.Errorf("неверный формат ввода")
	}

	if isRoman(matches[1]) && isRoman(matches[3]) {
		useRoman = true
		var err error
		a, err = romanToArabic(matches[1])
		if err != nil {
			return "", err
		}
		b, err = romanToArabic(matches[3])
		if err != nil {
			return "", err
		}
	} else if isArabic(matches[1]) && isArabic(matches[3]) {
		useRoman = false
		a, _ = strconv.Atoi(matches[1])
		b, _ = strconv.Atoi(matches[3])
	} else {
		return "", fmt.Errorf("разные системы счисления")
	}
	operation = matches[2]

	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			return "", fmt.Errorf("деление на ноль")
		}
		result = a / b
	default:
		return "", fmt.Errorf("неизвестная операция")
	}
	if useRoman {
		if result <= 0 {
			return "", fmt.Errorf("результат вне диапазона для римских чисел")
		}
		return arabicToRoman(result), nil
	}
	return fmt.Sprintf("%d", result), nil
}

func main() {
	fmt.Print("Введите операцию (например, 5 + 3 или V + II): ")
	var input string
	fmt.Scanln(&input)

	result, err := calculator(input)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Результат:", result)
}
