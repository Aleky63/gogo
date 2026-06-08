package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var (
	mtx   = sync.Mutex{}
	money = 1000
	bank  = 0
)

func payHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("fail to read HTTP Body!", err.Error())
		return
	}
	httpRequestBodyString := string(httpRequestBody)
	paymentAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		fmt.Println("fail to read HTTP body  to intereger!", err.Error())
		return
	}

	mtx.Lock()
	if money-paymentAmount >= 0 {

		money -= -paymentAmount
		fmt.Println("ОПЛАТА УСПЕШНО ПРОШЛА:", money)
	} else {
		fmt.Println("ОПЛАТА не прошла не хватает денег")
	}
	mtx.Unlock()
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("fail to read HTTP requet body!", err.Error())
		return
	}
	httpRequestBodyString := string(httpRequestBody)
	saveAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		fmt.Println("fail to convert HTTP body  to intereger!", err.Error())
		return

	}
	mtx.Lock()
	if money >= saveAmount {
		money -= saveAmount
		bank += saveAmount
		fmt.Println("Деньги УСПЕШНО прошли в банк:", money)
		fmt.Println("Новое значение денег на счете", bank)
	} else {
		fmt.Println("Не хватает денег положить на счет.")
	}
	mtx.Unlock()
}

func main() {
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHandler)

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("HTTP server error!", err.Error())
	}

	fmt.Println("hello")
}
