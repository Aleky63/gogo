package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/fatih/color"
)

type Payment struct {
	Description string `json:"description"`
	USD         int    `json:"usd"`
	FullName    string `json:"fullName"`
	Address     string `json:"address"`
	Array       []int  `json:"array"`

	Bool bool `json:"bool"`
	Time time.Time
}

func (p Payment) Println() {
	magenta := color.New(color.FgHiMagenta).SprintFunc()
	fmt.Println(magenta("================="))

	fmt.Println("Description:", p.Description)
	fmt.Println("USD:", p.USD)
	fmt.Println("FullName:", p.FullName)
	fmt.Println("Address:", p.Address)
	fmt.Println("Array:", p.Array)
	fmt.Println("Bool:", p.Bool)

	fmt.Println(magenta("================="))
}

var (
	mtx            = sync.Mutex{}
	money          = 1000
	paymentHistory = make([]Payment, 0)
)

type HttpResponse struct {
	Money          int       `json:"mmmoney"`
	PaymentHistory []Payment `json:"phistory"`
}

func payHandler(w http.ResponseWriter, r *http.Request) {
	var payment Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	payment.Time = time.Now()
	payment.Println()

	mtx.Lock()
	defer mtx.Unlock()
	if money-payment.USD >= 0 {
		money -= payment.USD
	}

	paymentHistory = append(paymentHistory, payment)

	HttpResponse := HttpResponse{
		Money: money,

		PaymentHistory: paymentHistory,
	}

	b, err := json.MarshalIndent(HttpResponse, "", "	")
	if err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("err:", err)
		return
	}
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fooParam := r.URL.Query().Get("foo")
	booParam := r.URL.Query().Get("boo")
	fmt.Println("Получены параметры:", fooParam, booParam)
	response := map[string]interface{}{
		"foo":     fooParam,
		"boo":     booParam,
		"message": "Обработчик /default вызван",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/default", defaultHandler)

	fmt.Println("Сервер запущен на порту :9091")
	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("Ошибка запуска HTTP сервера:", err)
	}
}
