package main

import (
	"fmt"
	"net/http"
	"sync"
)

type Payment struct {
	Description string
	USD         int
}

var (
	mtx            = sync.Mutex{}
	money          = 1000
	paymentHistory = make([]Payment, 0)
)

func payHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/pay", payHandler)

	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("Ошибка во время работы HTTP сервера:", err)
	}
}
