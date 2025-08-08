package main

import (
	"fmt"
	"slices"
	"strings"
)

func countFriends(data map[string][]string) map[string]int {
	friendsCount := make(map[string]int, len(data))
	names := make([]string, 0, len(data))

	for name, friends := range data {
		friendsCount[name] = len(friends)
		names = append(names, name)
	}
	slices.Sort(names)

	fmt.Println("Количество друзей:")
	for _, name := range names {
		fmt.Printf("%s: %d\n", name, friendsCount[name])
	}
	return friendsCount
}

func commonFriends(data map[string][]string, user1, user2 string) []string {
	friends1, ok1 := data[user1]
	friends2, ok2 := data[user2]
	if !ok1 || !ok2 {
		fmt.Printf("Один из пользователей (%s или %s) не найден\n", user1, user2)
		return nil
	}

	common := make([]string, 0, min(len(friends1), len(friends2)))
	for _, friend := range friends1 {
		if slices.Contains(friends2, friend) {
			common = append(common, friend)
		}
	}
	slices.Sort(common)
	return common
}

func mostPopularUsers(data map[string][]string) ([]string, int) {
	maxFriends := 0
	popularUsers := make([]string, 0)

	
	for _, friends := range data {
		if len(friends) > maxFriends {
			maxFriends = len(friends)
		}
	}

	
	for user, friends := range data {
		if len(friends) == maxFriends {
			popularUsers = append(popularUsers, user)
		}
	}

	slices.Sort(popularUsers)
	return popularUsers, maxFriends
}

func main() {
	friendsData := map[string][]string{
		"Алексей":  {"Иван", "Сергей", "Елена"},
		"Иван":     {"Алексей", "Дмитрий", "Мария"},
		"Сергей":   {"Алексей", "Елена"},
		"Дмитрий":  {"Иван", "Елена", "Ольга"},
		"Елена":    {"Алексей", "Сергей", "Дмитрий"},
		"Мария":    {"Иван", "Ольга"},
		"Ольга":    {"Дмитрий", "Мария"},
		"Анна":     {"Петр"},
		"Петр":     {"Анна", "Сергей"},
		"Светлана": {"Иван", "Елена"},
	}

	countFriends(friendsData)

	user1, user2 := "Иван", "Елена"
	common := commonFriends(friendsData, user1, user2)
	if len(common) > 0 {
		fmt.Printf("Общие друзья между %s и %s: %s\n", 
			user1, user2, strings.Join(common, ", "))
	} else {
		fmt.Printf("У %s и %s нет общих друзей\n", user1, user2)
	}

	
	popular, max := mostPopularUsers(friendsData)
	fmt.Printf("Самые популярные пользователи: %s (друзей: %d)\n", 
		strings.Join(popular, ", "), max)
}

//https://stepik.org/lesson/1535976/step/1?unit=1556545