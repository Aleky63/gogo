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
	for k, v := range r.Header {
		fmt.Println("k:", k, "--v:", v)
	}
	fmt.Println("Http method:", r.Method)

	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := "fail to read HTTP Body!" + err.Error()
		fmt.Println(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response", err)
		}
		return
	}
	httpRequestBodyString := string(httpRequestBody)

	paymentAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := "fail to convert HTTP body to integer!" + err.Error()
		fmt.Println(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response", err)
		}
		return
	}

	mtx.Lock()
	defer mtx.Unlock()

	if money >= paymentAmount {
		money -= paymentAmount
		msg := "ОПЛАТА УСПЕШНО ПРОШЛА! Остаток на счете: " + strconv.Itoa(money)
		fmt.Println(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response", err)
		}
	} else {
		msg := "Недостаточно средств! Доступно: " + strconv.Itoa(money)
		fmt.Println(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response", err)
		}
	}
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		msg := "fail to read HTTP request body!" + err.Error()
		fmt.Println(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response", err)
		}
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	saveAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		msg := "fail to convert HTTP body to integer!" + err.Error()
		fmt.Println(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response", err)
		}
		return
	}

	mtx.Lock()
	defer mtx.Unlock()

	if money >= saveAmount {
		// Переводим деньги со счета в банк
		money -= saveAmount
		bank += saveAmount

		msg := "ДЕНЬГИ ПЕРЕВЕДЕНЫ В БАНК! Остаток на счете: " + strconv.Itoa(money)
		fmt.Println(msg)

		msgbank := "Накопления в банке: " + strconv.Itoa(bank)
		fmt.Println(msgbank)

		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response", err)
		}
		_, err = w.Write([]byte("\n" + msgbank))
		if err != nil {
			fmt.Println("fail to write HTTP response", err)
		}
	} else {
		msg := "Недостаточно денег на счете для перевода! Доступно: " + strconv.Itoa(money)
		fmt.Println(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response", err)
		}
	}
}

func main() {
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHandler)

	fmt.Println("========================================")
	fmt.Println("🚀 Сервер запущен на http://localhost:9091")
	fmt.Println("========================================")
	fmt.Println("Доступные операции:")
	fmt.Println("  POST /pay  - списать деньги (уменьшает money)")
	fmt.Println("  POST /save - перевести деньги в банк (money -> bank)")
	fmt.Println("========================================")
	fmt.Println("Примеры запросов:")
	fmt.Println("  curl -X POST -d \"500\" http://localhost:9091/save")
	fmt.Println("  curl -X POST -d \"300\" http://localhost:9091/pay")
	fmt.Println("========================================")

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("❌ HTTP server error!", err.Error())
	}
}
