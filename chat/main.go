package main

import (
	"errors"
	"fmt"
)

// Данные
type Message struct {
	From    string
	Message string
}

type Chat struct {
	ID       int
	Messages []Message
}

// Ошибка базы данных
type DatabaseError struct {
	Message string
	Code    int
}

func (e DatabaseError) Error() string {
	return fmt.Sprintf("database error with code %d: %s", e.Code, e.Message)
}

// Сервис для работы
type Service struct {
	db interface {
		GetChatByIDWithMessages(id int) (*Chat, error)
	}
}

func (s Service) PrintChat(id int) {
	var dbErr *DatabaseError

	chat, err := s.db.GetChatByIDWithMessages(id)
	if errors.As(err, &dbErr) {
		if dbErr.Code == 24 {
			fmt.Printf("Ошибка запроса: %s", dbErr)
		} else {
			fmt.Printf("Инфраструктурная ошибка: %s", dbErr)
		}
		return
	} else if err != nil {
		fmt.Printf("Неизвестная ошибка: %s", err)
		return
	}
	for _, message := range chat.Messages {
		fmt.Printf("%s: %s\n", message.From, message.Message)
	}
}

// Мок базы данных для тестирования
type MockDB struct{}

func (m MockDB) GetChatByIDWithMessages(id int) (*Chat, error) {
	switch id {
	case 1:
		return &Chat{
			ID: 1,
			Messages: []Message{
				{From: "Alice", Message: "Привет!"},
				{From: "Bob", Message: "Здравствуйте!"},
				{From: "Alice", Message: "Как дела?"},
			},
		}, nil
	case 2:
		return nil, DatabaseError{Message: "запрос не прошёл", Code: 24} // Ошибка запроса
	case 3:
		return nil, DatabaseError{Message: "ошибка подключения", Code: 500} // Инфраструктурная ошибка
	case 4:
		return nil, &DatabaseError{Message: "ошибка парсинга", Code: 24} // Указатель и код 24
	default:
		return nil, fmt.Errorf("неизвестная ошибка")
	}
}

func main() {
	service := Service{db: MockDB{}}

	fmt.Println("=== Чат 1 (успех) ===")
	service.PrintChat(1)

	fmt.Println("\n=== Чат 2 (ошибка запроса, DatabaseError) ===")
	service.PrintChat(2)

	fmt.Println("\n=== Чат 3 (инфраструктурная ошибка) ===")
	service.PrintChat(3)

	fmt.Println("\n=== Чат 4 (ошибка с указателем, код 24) ===")
	service.PrintChat(4)

	fmt.Println("\n=== Чат 5 (неизвестная ошибка) ===")
	service.PrintChat(5)

}

// https://stepik.org/lesson/1501024/step/4?unit=1521140
