package main

import "fmt"

type Order struct {
	ID       int
	Quantity int
}

func supplier(ordersChan chan<- Order) {
	for i := 1; i <= 5; i++ {
		order := Order{
			ID:       i,
			Quantity: i * 10,
		}
		ordersChan <- order
	}

	close(ordersChan)
}

func warehouse(ordersChan <-chan Order, resultsChan chan<- string) {

	stock := 101

	for order := range ordersChan {
		if order.Quantity <= stock {
			// Товара достаточно — обрабатываем заказ
			stock -= order.Quantity
			result := fmt.Sprintf("Order %d processed, left %d items", order.ID, stock)
			resultsChan <- result
		} else {
			// Товара не хватает — отклоняем заказ
			result := fmt.Sprintf("Order %d rejected: not enough items", order.ID)
			resultsChan <- result
		}
	}

	// Все заказы обработаны — закрываем канал результатов
	close(resultsChan)
}

func main() {
	ordersChan := make(chan Order)
	resultsChan := make(chan string)

	go supplier(ordersChan)
	go warehouse(ordersChan, resultsChan)

	for msg := range resultsChan {
		fmt.Println(msg)
	}
	fmt.Println("\u2705 Все заказы обработаны. Программа завершена.❤️")
}
