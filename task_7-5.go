package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Игра 'Угадай число'")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		quit := gamePlayLogic()

		if quit {
			break
		}

		fmt.Print("Еще раз играем? да/нет: ")
		scanner.Scan()
		answer := strings.TrimSpace(scanner.Text())

		if strings.ToLower(answer) == "нет" {
			break
		} else if strings.ToLower(answer) == "да" {
			gamePlayLogic()
		} else {
			fmt.Println("Некорректный ввод")
		}
	}

	fmt.Println("ББ")
}

func generateRandom(min, max int) int {
	n := rand.Intn(max-min+1) + min
	return n
}

func readChel() (int, bool, error) {
	fmt.Print("Введите вариант ответа: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()
	input = strings.TrimSpace(input)

	if strings.ToLower(input) == "выход" {
		return 0, true, nil
	}

	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, false, err
	}
	return num, false, nil
}

func gamePlayLogic() bool {
	numRand := generateRandom(1, 100)
	attempts := 0

	for {
		num, quit, err := readChel()
		if quit {
			fmt.Println("Выход из игры")
			break
		}
		if err != nil {
			fmt.Println("Ошибка ввода")
			continue
		}

		attempts++

		if num < numRand {
			fmt.Println("Загаданное число больше")
		} else if num > numRand {
			fmt.Println("Загаданное число меньше")
		} else {
			fmt.Printf("Вы угадали число - %d с %d попытки\n", numRand, attempts)
			break
		}
	}
	return false
}
