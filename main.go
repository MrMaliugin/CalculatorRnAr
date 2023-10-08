package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var RN = map[string]int{
	"I":        1,
	"II":       2,
	"III":      3,
	"IV":       4,
	"V":        5,
	"VI":       6,
	"VII":      7,
	"VIII":     8,
	"IX":       9,
	"X":        10,
	"XI":       11,
	"XII":      12,
	"XIII":     13,
	"XIV":      14,
	"XV":       15,
	"XVI":      16,
	"XVII":     17,
	"XVIII":    18,
	"XIX":      19,
	"XX":       20,
	"XXI":      21,
	"XXII":     22,
	"XXIII":    23,
	"XXIV":     24,
	"XXV":      25,
	"XXVI":     26,
	"XXVII":    27,
	"XXVIII":   28,
	"XXIX":     29,
	"XXX":      30,
	"XXXI":     31,
	"XXXII":    32,
	"XXXIII":   33,
	"XXXIV":    34,
	"XXXV":     35,
	"XXXVI":    36,
	"XXXVII":   37,
	"XXXVIII":  38,
	"XXXIX":    39,
	"XL":       40,
	"XLI":      41,
	"XLII":     42,
	"XLIII":    43,
	"XLIV":     44,
	"XLV":      45,
	"XLVI":     46,
	"XLVII":    47,
	"XLVIII":   48,
	"XLIX":     49,
	"L":        50,
	"LI":       51,
	"LII":      52,
	"LIII":     53,
	"LIV":      54,
	"LV":       55,
	"LVI":      56,
	"LVII":     57,
	"LVIII":    58,
	"LIX":      59,
	"LX":       60,
	"LXI":      61,
	"LXII":     62,
	"LXIII":    63,
	"LXIV":     64,
	"LXV":      65,
	"LXVI":     66,
	"LXVII":    67,
	"LXVIII":   68,
	"LXIX":     69,
	"LXX":      70,
	"LXXI":     71,
	"LXXII":    72,
	"LXXIII":   73,
	"LXXIV":    74,
	"LXXV":     75,
	"LXXVI":    76,
	"LXXVII":   77,
	"LXXVIII":  78,
	"LXXIX":    79,
	"LXXX":     80,
	"LXXXI":    81,
	"LXXXII":   82,
	"LXXXIII":  83,
	"LXXXIV":   84,
	"LXXXV":    85,
	"LXXXVI":   86,
	"LXXXVII":  87,
	"LXXXVIII": 88,
	"LXXXIX":   89,
	"XC":       90,
	"XCI":      91,
	"XCII":     92,
	"XCIII":    93,
	"XCIV":     94,
	"XCV":      95,
	"XCVI":     96,
	"XCVII":    97,
	"XCVIII":   98,
	"XCIX":     99,
	"C":        100,
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	expression, _ := reader.ReadString('\n')
	expression = strings.TrimSpace(expression)
	elements := strings.Split(expression, " ")

	for len(elements) == 1 {
		fmt.Println("Ошибка: строка не является математической операцией.")
		return
	}

	for len(elements) != 3 {
		fmt.Println("Ошибка: неверный формат выражения")
		return
	}

	operand1, operator, operand2 := elements[0], elements[1], elements[2]

	if Roman(operand1) && Roman(operand2) {
		CalculatorRN(operand1, operator, operand2)
	} else if Arabic(operand1) && Arabic(operand2) {
		CalculationAr(operand1, operator, operand2)
	} else {
		fmt.Println("Ошибка: неверный формат операндов")
	}
}

func Roman(s string) bool {

	for _, r := range s {
		if r != 'I' && r != 'V' && r != 'X' {
			return false
		}
	}
	return true
}

func Arabic(s string) bool {

	_, err := strconv.Atoi(s)
	return err == nil
}

func CalculatorRN(operand1, operator, operand2 string) {

	num1, num2 := RN[operand1], RN[operand2]

	if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
		fmt.Println("Используйте числа от I до X")
		return
	}

	switch operator {
	case "+":
		result := num1 + num2
		fmt.Println(convertRN(result))
	case "-":
		result := num1 - num2
		if result < 0 {
			fmt.Println("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
		} else {
			fmt.Println(convertRN(result))
		}
	case "*":
		result := num1 * num2
		fmt.Println(convertRN(result))
	case "/":
		result := num1 / num2
		if result == 0 {
			fmt.Println("I")
		} else {
			fmt.Println(convertRN(result))
		}
	default:
		fmt.Println("Ошибка: неверный оператор")
	}
}

func CalculationAr(operand1, operator, operand2 string) {

	num1, err := strconv.Atoi(operand1)
	if err != nil {
		log.Fatal(err)
	}

	num2, err := strconv.Atoi(operand2)
	if err != nil {
		log.Fatal(err)
	}

	if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
		fmt.Println("Используйте числа от 1 до 10")
		return
	}

	switch operator {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		fmt.Println(num1 / num2)
	default:
		fmt.Println("Ошибка: неверный оператор")
	}
}

func convertRN(value int) string {

	for k, v := range RN {
		if v == value {
			return k
		}
	}
	return ""
}
