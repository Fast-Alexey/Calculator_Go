package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var err error

func main() {
	reader := bufio.NewReader(os.Stdin)
	var sL []string
	for {
		fmt.Println("Введите выражение для расчета")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		sL = strings.Split(text, " ")
		break
	}
	if len(sL) > 3 || len(sL) < 3 {
		err = fmt.Errorf("Вы ввели не корректные данные, программа завершена")
	}
	if err != nil {
		panic(err)
	}
	numS1 := sL[0]
	sign := sL[1]
	numS2 := sL[2]
	number1, _ := strconv.Atoi(numS1)
	number2, _ := strconv.Atoi(numS2)
	if number1 == 0 || number2 == 0 {
		fmt.Printf("Ответ = %v\n", TranslIntoArabic(numS1, numS2, sign))
	}
	if number1 < 0 || number2 < 0 || number1 > 10 || number2 > 10 {
		err = fmt.Errorf("Вы ввели не корректные данные, программа завершена")
	}
	if err != nil {
		panic(err)
	}
	if number1 > 0 || number2 > 0 {
		fmt.Printf("Ответ = %v\n", Account(number1, number2, sign))
	}
}

func TranslIntoArabic(numS1, numS2, sign string) string {
	romeToArab := map[string]int{
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
	numI1 := romeToArab[numS1]
	numI2 := romeToArab[numS2]
	if numI1 == 0 || numI2 == 0 {
		err = fmt.Errorf("Вы ввели не корректные данные, программа завершена")
	}
	if err != nil {
		panic(err)
	}
	if sign == "+" {
		ad := numI1 + numI2
		return transf(ad)
	}
	if sign == "-" {
		ad := numI1 - numI2
		return transf(ad)
	}
	if sign == "*" {
		ad := numI1 * numI2
		return transf(ad)
	}
	if sign == "/" {
		ad := numI1 / numI2
		return transf(ad)
	} else {
		err = fmt.Errorf("Вы ввели не корректные данные, программа завершена")
		panic(err)
	}
}

func transf(ad int) string {
	arrToRome := map[string]string{
		"1":  "X",
		"2":  "XX",
		"3":  "XX",
		"4":  "XL",
		"5":  "L",
		"6":  "LX",
		"7":  "LXX",
		"8":  "LXXX",
		"9":  "XC",
		"10": "C",
	}
	arabToRome := map[int]string{
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
	if ad < 1 {
		err = fmt.Errorf("Вы ввели не корректные данные, программа завершена")
		panic(err)
	}
	a1 := ad / 10
	a2 := ad % 10
	if a2 == 0 {
		return arrToRome[strconv.Itoa(a1)]
	} else {
		return arrToRome[strconv.Itoa(a1)] + arabToRome[a2]
	}
}
func Account(number1, number2 int, sign string) int {
	if sign == "+" {
		return number1 + number2
	}
	if sign == "-" {
		return number1 - number2
	}
	if sign == "*" {
		return number1 * number2
	}
	if sign == "/" {
		return number1 / number2
	} else {
		err = fmt.Errorf("Вы ввели не корректные данные(Знак), программа завершена")
		panic(err)
	}
}
