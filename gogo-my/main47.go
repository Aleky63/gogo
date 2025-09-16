package main

import (
	"errors"
	"fmt"
	"strings"
)

type Address struct {
	ID int

	City       string
	Street     string
	PostalCode int
}

type Product struct {
	ID          int
	Name        string
	Description string
	Price       int
	Category    string
	Brand       string
	Rating      float64
	Reviews     int
}

type CartItem struct {
	Product  Product
	Quantity int
}

type Useras struct {
	Name    string
	Email   string
	Phone   string
	Address Address
	Cart    []CartItem
}

func printInfo(useras Useras) error {
	if useras.Name == "" || useras.Phone == "" || useras.Address.City == "" || useras.Address.Street == "" {
		return errors.New("ошибка: не все обязательные данные пользователя переданы")
	}
	fmt.Printf("Покупатель %s. Телефон: %s. Адрес: г. %s, %s.\n",
		useras.Name, useras.Phone, useras.Address.City, useras.Address.Street)
	return nil
}

func hasElectronics(useras Useras) bool {
	for _, item := range useras.Cart {
		if item.Product.Category == "Электроника" {
			return true
		}
	}
	return false
}

func expensiveItems(useras Useras) []string {
	var res []string
	for _, item := range useras.Cart {
		if item.Product.Price >= 10000 {
			res = append(res, item.Product.Name)
		}

	}
	return res
}
func totalSum(useras Useras) int {
	var sum int
	for _, item := range useras.Cart {
		sum += item.Product.Price * int(item.Quantity)
	}
	return sum
}

func main47() {
	useras := Useras{
		Name:  "Иван Петров",
		Phone: "+7 999 123-45-67",
		Email: "ssss.ivanov@example.com",
		Address: Address{
			ID:         1,
			City:       "Yeysk",
			Street:     "Улица Высоцкого",
			PostalCode: 101000,
		},
		Cart: []CartItem{
			{
				Product: Product{
					ID:          1,
					Name:        "Ноутбук",
					Description: "Мощный ноутбук для работы и игр",
					Price:       59990,
					Category:    "Электроника",
					Brand:       "Brand A",
					Rating:      4.5,
					Reviews:     120,
				},
				Quantity: 1,
			},
			{
				Product: Product{
					ID:          2,
					Name:        "Смартфон",
					Description: "Современный смартфон с отличной камерой",
					Price:       29990,
					Category:    "Электроника",
					Brand:       "Brand B",
					Rating:      4.7,
					Reviews:     200,
				},
				Quantity: 2,
			},
			{
				Product: Product{
					ID:          3,
					Name:        "Наушники",
					Description: "Беспроводные наушники с шумоподавлением",
					Price:       1990,
					Category:    "Аудио",
					Brand:       "Brand C",
					Rating:      4.3,
					Reviews:     80,
				},
				Quantity: 1,
			},
		},
	}

	if err := printInfo(useras); err != nil {
		fmt.Println(err)
		return
	}

	if hasElectronics(useras) {
		fmt.Println("Пользователь является покупателем электроники.")
	} else {
		fmt.Println("Пользователь не является покупателем электроники.")
	}
	expItems := expensiveItems(useras)
	fmt.Printf("Товары в корзине, где цена 10000 и более: [%s].\n", strings.Join(expItems, ", "))
	fmt.Printf("Общая сумма покупки: %.2d руб.\n", totalSum(useras))
}

// https://stepik.org/lesson/1500867/step/3?unit=1520984
