package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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
var money = 1000
var paymenHistory = make([]Payment, 0)

func payHandler(w http.ResponseWriter, r *http.Request) {
	var payment Payment

	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	payment.Println()

	mtx.Lock()
	if money-payment.USD >= 0 {
		money -= payment.USD
	}

	paymenHistory = append(paymenHistory, payment)

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
