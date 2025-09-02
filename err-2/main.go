package main

import (
	"fmt"
	"log"
)

type MyError struct {
	Code    int
	Message string
}

// Error implements error.
func (m MyError) Error() string {

	return fmt.Sprintf("Ошибка %d: %s", m.Code, m.Message)
}

func someFunc() error {
	return &MyError{
		Code:    500,
		Message: "My error today!!!!!",
	}
}

func main() {
	if err := someFunc(); err != nil {
		log.Fatalf("ОШИБКА: %s", err)

	}
}

// https://stepik.org/lesson/1501023/step/1?unit=1521139
