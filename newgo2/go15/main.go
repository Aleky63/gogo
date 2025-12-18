package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"

	"github.com/fatih/color"
)

var mtx = sync.Mutex{}
var money = 1000
var bank = 0

func payHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Println("k:", k, "---v:", v)
	}

	defer r.Body.Close()
	httpRequestBody, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := "🤸🏾fail to read HTTP body:" + err.Error()
		fmt.Println(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("🤸🏾fail to write HTTP response:", err)
		}
		return

	}
	httpRequestBodyString := string(httpRequestBody)

	paymentAmount, err := strconv.Atoi(httpRequestBodyString)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := "🤸🏾 fail to convert HTTP body to integer:" + err.Error()
		fmt.Println(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("🤸🏾fail to write HTTP response:", err)
		}
		return
	}

	mtx.Lock()
	defer mtx.Unlock()

	if money-paymentAmount < 0 {
		w.WriteHeader(http.StatusBadRequest)
		msg := "❌ not enough money"
		fmt.Println(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("🤸🏾fail to write HTTP response:", err)
		}
		return
	}
	money -= paymentAmount
	msg := "👌👌👌Pay good. Balance: " + strconv.Itoa(money)
	fmt.Println(msg)
	_, err = w.Write([]byte(msg))
	if err != nil {
		fmt.Println("🤸🏾fail to write HTTP response:", err)
	}
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	httpRequestBody, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		msg := "🤸🏾fail to read HTTP body:" + err.Error()
		fmt.Println(msg)

		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("🤸🏾fail to write HTTP response:", err)
		}

		return
	}
	httpRequestBodyString := string(httpRequestBody)

	saveAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := "🤸🏾 fail to convert HTTP body to integer:" + err.Error()
		fmt.Println(msg)
		_, err = w.Write([]byte(msg))

		if err != nil {
			fmt.Println("🤸🏾fail to write HTTP response:", err)
		}
		return
	}
	mtx.Lock()
	defer mtx.Unlock()

	if money >= saveAmount {
		money -= saveAmount
		bank += saveAmount
		fmt.Println("👌👌👌Pay good. New summa money:", money)
		fmt.Println("👌👌👌Pay good. New summa bank:", bank)

	}

}

func main() {

	red := color.New(color.FgHiRed).SprintFunc()
	fmt.Println(red("-----🚕🚓🚕-----"))

	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHandler)

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println(" HTTP server ERRoOR:", err.Error())

	}
	fmt.Println(red("ЗАКОНЧИЛА ПРОГРАММА ВЫПОЛНЕНИЕ"))
}
