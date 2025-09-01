package main

import "fmt"

type User struct {
	ID   string
	Name string
}

func UpdateUserName(id, name string) error {
	user, err := GetUserByID(id)
	if err != nil {
		return fmt.Errorf("failed to get user by id %s: %w", id, err)
	}

	user.Name = name

	if err := SaveUser(user); err != nil {
		return fmt.Errorf("failed to save user %s: %w", id, err)
	}

	return nil
}

func GetUserByID(id string) (*User, error) {
	// допустим нашли пользователя
	if id == "123" {
		return &User{ID: "123", Name: "OldName"}, nil
	}
	return nil, fmt.Errorf("user with id %s not found", id)
}

func SaveUser(user *User) error {
	// допустим всё сохранилось
	if user.ID == "" {
		return fmt.Errorf("cannot save user without ID")
	}
	return nil
}

func main() {
	if err := UpdateUserName("123", "NewName"); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Имя пользователя обновлено успешно")
	}
}
