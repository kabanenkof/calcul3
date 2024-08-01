package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func сalculate(expression string) {
	var (
		mathOperation byte
		totalData     []string
	)

	if strings.Contains(expression, " + ") {
		totalData = strings.Split(expression, " + ")
		mathOperation = '+'
	} else if strings.Contains(expression, " - ") {
		totalData = strings.Split(expression, " - ")
		mathOperation = '-'
	} else if strings.Contains(expression, " * ") {
		totalData = strings.Split(expression, " * ")
		mathOperation = '*'
	} else if strings.Contains(expression, " / ") {
		totalData = strings.Split(expression, " / ")
		mathOperation = '/'
	} else {
		panic("Incorrect math operation")
	}

	if mathOperation == '*' || mathOperation == '/' {
		if strings.Contains(totalData[1], "\"") {
			panic("String should divide or multiply by number. not a string!")
		}
	}

	for i := 0; i < len(totalData); i++ {
		totalData[i] = strings.Replace(totalData[i], "\"", "", -1)
	}

	switch mathOperation {
	case '+':
		answer := totalData[0] + totalData[1]
		printResultInQuotes(checkAnswerLength(answer))
	case '-':
		index := strings.Index(totalData[0], totalData[1])
		if index == -1 {
			printResultInQuotes(totalData[0])
		} else {
			answer := totalData[0][:index] + totalData[0][index+len(totalData[1]):]
			printResultInQuotes(checkAnswerLength(answer))
		}
	case '*':
		multiplier, _ := strconv.Atoi(totalData[1])
		var answer string
		for i := 0; i < multiplier; i++ {
			answer += totalData[0]
		}
		printResultInQuotes(checkAnswerLength(answer))
	case '/':
		firstStrLen := len(totalData[0])
		secondStrLen, _ := strconv.Atoi(totalData[1])
		newStringLen := firstStrLen / secondStrLen
		answer := totalData[0][:newStringLen]
		printResultInQuotes(checkAnswerLength(answer))

	default:
		panic("Ooops! Smthng wrong!")
	}
}

func printResultInQuotes(text string) {
	fmt.Println("\"" + text + "\"")
}
func checkAnswerLength(answer string) string {
	currentLength := len(answer)

	if currentLength >= 11 {
		panic("Ошибка: строка превышает 10 символов.")
	}

	correctAnswer := ""
	if currentLength > 40 {
		correctAnswer = answer[:40] + "..."
	} else {
		correctAnswer = answer
	}

	return correctAnswer
}
func main() {
	inputScanner := bufio.NewScanner(os.Stdin)
	inputScanner.Scan()
	expression := inputScanner.Text()

	сalculate(expression)
}
