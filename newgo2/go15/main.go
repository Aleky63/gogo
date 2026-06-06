package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync/atomic"
)

var (
	money = atomic.Int64{}
	bank  = atomic.Int64{}
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
	if money.Load()-int64(paymentAmount) >= 0 {
		money.Add(int64(-paymentAmount))
		fmt.Println("ОПЛАТА УСПЕШНО ПРОШЛА:", money.Load())
	} else {
		fmt.Println("ОПЛАТА не прошла не хватает денег")
	}
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
	if money.Load() >= int64(saveAmount) {
		money.Add(int64(-saveAmount))
		bank.Add(int64(saveAmount))
		fmt.Println("Деньги УСПЕШНО прошли в банк:", money.Load())
		fmt.Println("Новое значение денег на счете", bank.Load())
	} else {
		fmt.Println("Не хватает денег положить на счет.")
	}
}

func main() {
	money.Add(1000)
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHandler)

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("HTTP server error!", err.Error())
	}

	fmt.Println("hello")
}
