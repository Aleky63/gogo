package main

import (
	"errors"
	"fmt"

	"strings"
)

const (
	MinAge                = 15
	MaxAge                = 80
	MinGrade              = 1
	MaxGrade              = 5
	GreatAge              = 30
	MinGradeAfterGreatAge = 3
)

var (
	ErrEmptyName         = errors.New("name cannot be empty")
	ErrTooYoung          = errors.New("too young")
	ErrTooOld            = errors.New("too old")
	ErrGradeOutOfRange   = errors.New("grade out of range")
	ErrTooLowGradeForAge = errors.New("too low grade for age")
	ErrIncorrectEmail    = errors.New("incorrect email")
)

type Student1 struct {
	Name  string
	Age   int
	Grade int
	Email string
}

func NewStudent1(name string, age, grade int, email string) (*Student1, error) {

	if strings.TrimSpace(name) == "" {
		return nil, ErrEmptyName
	}

	if age < MinAge {
		return nil, ErrTooYoung
	}
	if age > MaxAge {
		return nil, ErrTooOld
	}

	if grade < MinGrade || grade > MaxGrade {
		return nil, ErrGradeOutOfRange
	}

	if age > GreatAge && grade < MinGradeAfterGreatAge {
		return nil, ErrTooLowGradeForAge
	}

	if !strings.Contains(email, "@") {
		return nil, ErrIncorrectEmail
	}

	return &Student1{
		Name:  name,
		Age:   age,
		Grade: grade,
		Email: email,
	}, nil
}

// Пример использования
func main50() {
	// Тестовые случаи
	testCases := []struct {
		name  string
		age   int
		grade int
		email string
		desc  string
	}{
		{"Иван Петров", 20, 4, "ivan@example.com", "Валидный студент"},
		{"", 20, 4, "test@example.com", "Пустое имя"},
		{"Анна Смирнова", 14, 4, "anna@example.com", "Слишком молодой"},
		{"Петр Иванов", 85, 4, "petr@example.com", "Слишком старый"},
		{"Мария Козлова", 25, 6, "maria@example.com", "Оценка вне диапазона"},
		{"Алексей Сидоров", 35, 2, "alex@example.com", "Низкая оценка для возраста"},
		{"Елена Волкова", 22, 3, "elena.example.com", "Неправильный email"},
		{"Дмитрий Новиков", 32, 4, "dmitry@example.com", "Валидный зрелый студент"},
	}

	for i, tc := range testCases {
		Student1, err := NewStudent1(tc.name, tc.age, tc.grade, tc.email)
		fmt.Printf("Тест %d (%s):\n", i+1, tc.desc)
		if err != nil {
			fmt.Printf("  Ошибка: %v\n", err)
		} else {
			fmt.Printf("  Студент создан: %+v\n", *Student1)
		}
		fmt.Println()
	}
}

// https://stepik.org/lesson/1500872/step/5?thread=solutions&unit=1520989
