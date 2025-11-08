package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	text, err := GetInput()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	letters, digits, spaces, punctuation := CountCharacters(text)
	DisplayResults(letters, digits, spaces, punctuation)
}

func GetInput() (string, error) {
	fmt.Print("Введите ваш текст: ")

	scanner := bufio.NewScanner(os.Stdin)
	ok := scanner.Scan()
	if !ok {
		return "", fmt.Errorf("ошибка ввода текста")
	}

	text := scanner.Text()
	text = strings.TrimSpace(text)
	if text == "" {
		return "", fmt.Errorf("введен пустой текст")
	}

	return text, nil
}

func CountCharacters(text string) (letters, digits, spaces, punctuation int) {
	for _, r := range text {
		switch {
		case unicode.IsLetter(r):
			letters++
		case unicode.IsDigit(r):
			digits++
		case unicode.IsSpace(r):
			spaces++
		case unicode.IsPunct(r):
			punctuation++
		}
	}
	return
}

func DisplayResults(letters, digits, spaces, punctuation int) {
	fmt.Println("Количество букв:", letters)
	fmt.Println("Количество цифр:", digits)
	fmt.Println("Количество пробелов:", spaces)
	fmt.Println("Количество знаков препинания:", punctuation)
}
