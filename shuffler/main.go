// main.go
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"shuffler/questions"

	"github.com/fatih/color"
)

func main() {
	qs := questions.All()

	rand.Shuffle(len(qs), func(i, j int) {
		qs[i], qs[j] = qs[j], qs[i]
	})

	color.Cyan("🌍 Викторина: Угадай столицу!\n")
	fmt.Println("Введите 'выход', чтобы завершить.\n")

	scanner := bufio.NewScanner(os.Stdin)
	correct := 0

	for i, q := range qs {
		color.Yellow("Вопрос %d: Какая столица у %s?", i+1, q.Country)
		fmt.Print("Ваш ответ: ")

		if !scanner.Scan() {
			break
		}

		answer := strings.TrimSpace(scanner.Text())
		if strings.ToLower(answer) == "выход" {
			color.Red("\nВы вышли из игры.")
			break
		}

		// Сравниваем без учёта регистра и лишних пробелов
		if strings.EqualFold(answer, q.Capital) {
			color.Green("✅ Верно! Это %s.", q.Capital)
			correct++
		} else {
			color.Red("❌ Неверно. Правильный ответ: %s.", q.Capital)
		}
		fmt.Println()
	}

	total := len(qs)
	color.Cyan("\nИгра завершена! Ваш результат: %d из %d", correct, total)

	// Оценка (исправлено: явное преобразование типов)
	switch {
	case correct == total:
		color.Green("🏆 Идеально! Вы — географический гений!")
	case correct >= int(float64(total)*0.8):
		color.Green("👍 Отлично!")
	case correct >= int(float64(total)*0.6):
		color.Yellow("🙂 Неплохо!")
	default:
		color.Red("😅 Пора повторить карту мира!")
	}
}
