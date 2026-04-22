package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/fatih/color"
)

var mtx = sync.Mutex{}
var money = atomic.Int64{}
var bank = atomic.Int64{}

func payHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println(" fail to read HTTP body:", err.Error())
		return
	}
	httpRequestBodyString := string(httpRequestBody)

	paymentAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		fmt.Println(" fail to convert HTTP body to integer:", err)
		return
	}

	mtx.Lock()
	if money.Load()-int64(paymentAmount) >= 0 {
		time.Sleep((1 * time.Second))
		money.Add(-int64(paymentAmount))
		fmt.Println("👌👌👌Pay good:", money.Load())
	} else {
		fmt.Println(" 🛠️ Not enough money to buy:", money.Load())
	}
	mtx.Unlock()
}
func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println(" fail to read HTTP body:", err)
		return
	}
	httpRequestBodyString := string(httpRequestBody)

	saveAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		fmt.Println(" fail to convert HTTP body to integer:", err)
		return
	}
	mtx.Lock()
	if money.Load() >= int64(saveAmount) {
		money.Add(int64(-saveAmount))
		bank.Add(int64(saveAmount))
		fmt.Println("👌👌👌Pay good. New summa money:", money.Load())
		fmt.Println("👌👌👌Pay good. New summa bank:", bank.Load())
	} else {
		fmt.Println(" 🛠️  Don't have enough money to put in a piggy bank:", money.Load())
	}
	mtx.Unlock()
}

func main() {
	money.Add(1000)
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
