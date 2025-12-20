package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	// "strconv"
	// "strings"
	"sync"

	"github.com/fatih/color"
)

type Payment struct {
	Description string `json:"description"`
	USD         int    `json:"usd"`
	FullName    string `json:"fullName"`
	Address     string `json:"address"`
	Color       string `json:"color"`
	Flag        bool   `json:"flag"`
	Array       []int  `json:"array"`
	Time        time.Time
}

func (p Payment) Println() {
	fmt.Println("Description:", p.Description)
	fmt.Println("USD:", p.USD)
	fmt.Println("FullName:", p.FullName)
	fmt.Println("Address:", p.Address)
	fmt.Println("Color:", p.Color)
	fmt.Println("Flag:", p.Flag)
	fmt.Println("Array:", p.Array)
}

var mtx = sync.Mutex{}
var money = 10000
var paymenHistory = make([]Payment, 0)

type HttpResponse struct {
	Money          int
	PaymentHistory []Payment
}

func payHandler(w http.ResponseWriter, r *http.Request) {
   fooParam := r.URL.Query().Get("foo")
   booParam := r.URL.Query().Get("boo")

fmt.Println("fooParam:", fooParam)
fmt.Println("booParam:", booParam)

	var payment Payment

	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	payment.Time = time.Now()

	payment.Println()

	mtx.Lock()
	if money-payment.USD >= 0 {
		money -= payment.USD
	}

	paymenHistory = append(paymenHistory, payment)

	httpResponse := HttpResponse{
		Money:          money,
		PaymentHistory: paymenHistory,
	}

	b, err := json.Marshal(httpResponse)
	if err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(b); err != nil {
		fmt.Println("err:", err)
		return
	}

	red := color.New(color.FgHiRed).SprintFunc()
	yellow := color.New(color.FgHiYellow).SprintFunc()

	fmt.Println(red("money:", money))
	fmt.Println(yellow("payment history:", paymenHistory))

	mtx.Unlock()
}

func main() {

	fmt.Println("-----I🚓I-----")

	http.HandleFunc("/pay", payHandler)

	

	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("ERROR  during HTTP server operation:", err)
	}

	fmt.Println("")
}
