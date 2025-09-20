package main

import (
	"fmt"
)

// Симуляция склада товаров
var stock = map[int]int{
	1: 10, // ID 1: 10 единиц товара
	2: 0,  // ID 2: нет товара
	3: 5,  // ID 3: 5 единиц товара
}

// checkItemInStock проверяет, есть ли нужное количество товара на складе
func checkItemInStock(itemID int, quantity int) bool {
	return stock[itemID] >= quantity
}

// placeOrder оформляет заказ, проверяя наличие товара на складе
func placeOrder(itemID int, quantity int) error {
	if !checkItemInStock(itemID, quantity) {
		return fmt.Errorf("товара с ID %d нет в нужном количестве", itemID)
	}
	// Списываем товар со склада
	stock[itemID] -= quantity
	fmt.Printf("Заказ оформлен: ID товара %d, количество %d\n", itemID, quantity)
	return nil
}

func main() {
	// Попробуем оформить заказы
	if err := placeOrder(1, 5); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Остаток на складе:", stock[1])
	}

	if err := placeOrder(2, 1); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Остаток на складе:", stock[2])
	}

	if err := placeOrder(3, 6); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Остаток на складе:", stock[3])
	}
}
