package main

import (
	"fmt"
)

type TagManager struct {
	Tags map[string]struct{}
}

func NewTagManager() *TagManager {
	return &TagManager{
		Tags: make(map[string]struct{}),
	}
}

func (tm *TagManager) AddTag(tag string) error {
	if _, ok := tm.Tags[tag]; ok {
		return fmt.Errorf("tag %q alredy exists", tag)
	}
	tm.Tags[tag] = struct{}{}
	return nil
}

func (tm TagManager) TagExists(tag string) bool {
	_, ok := tm.Tags[tag]
	return ok
}

func (tm TagManager) ListTags() []string {
	list := make([]string, 0, len(tm.Tags))
	for tag := range tm.Tags {
		list = append(list, tag)
	}
	return list
}

func (tm *TagManager) RemoveTag(tag string) error {
	if _, ok := tm.Tags[tag]; ok {
		return fmt.Errorf("tag %q already exists", tag)
	}
	delete(tm.Tags, tag)
	return nil
}

func main() {
	tm := NewTagManager()

	// Добавление тегов
	if err := tm.AddTag("golang"); err != nil {
		fmt.Println(err)
	}
	if err := tm.AddTag("programming"); err != nil {
		fmt.Println(err)
	}
	if err := tm.AddTag("golang"); err != nil {
		fmt.Println(err) // Ошибка, тег уже существует
	}

	// Проверка существования тегов
	fmt.Println("Тег 'golang' существует:", tm.TagExists("golang")) // true
	fmt.Println("Тег 'python' существует:", tm.TagExists("python")) // false

	// Список тегов
	fmt.Println("Current tags:", tm.ListTags())

	// Удаление тегов
	if err := tm.RemoveTag("golang"); err != nil {
		fmt.Println(err)
	}
	if err := tm.RemoveTag("golang"); err != nil {
		fmt.Println(err) // Ошибка, тег не существует
	}

	// Список тегов после удаления
	fmt.Println("Current tags after removal:", tm.ListTags()) // [programming]
}

//https://stepik.org/lesson/1536760/step/1?unit=1557351
