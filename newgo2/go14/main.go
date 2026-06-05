package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fatih/color"
)

func payHandler(w http.ResponseWriter, r *http.Request) {
	str := "Buy Kyiv for money!!!"
	b := []byte(str)
	_, err := w.Write(b)
	if err != nil {
		fmt.Println("ERROOOOOOOR:", err.Error())
	} else {
		fmt.Println("BUY_________BUY")
	}
}

func cancelHandler(w http.ResponseWriter, r *http.Request) {
	str := "Kyiv in three days, didn't work, cancel!!!"
	b := []byte(str)
	_, err := w.Write(b)
	if err != nil {
		fmt.Println("ERROOOOOOOR:", err.Error())
	} else {
		fmt.Println("CANCEL______________CANCEL")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	str := "Kyiv in three days!!! "
	b := []byte(str)
	_, err := w.Write(b)
	if err != nil {
		fmt.Println("ERROOOOOOOR:", err.Error())
	} else {
		fmt.Println("OK_________🎈_____OK")
	}
}

func handlerProba(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello,Eysk!!!🦀🦀🦀"))
	if err != nil {
		fmt.Println("err:", err.Error())
	}
}

func handlerSleep(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	_, err := w.Write([]byte("HTTP respons!!!🐩🐈‍⬛🐩"))
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Println("Sleep🤦‍♀️🤦‍♀️🤦‍♀️🤦‍♀️🤦‍♀️🤦‍♀️🤦‍♀️🤦‍♀️🤦‍♀️🤦‍♀️🤦‍♀️")
	}
}

func main() {
	http.HandleFunc("/default", handler)
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/cancel", cancelHandler)

	http.HandleFunc("/", handlerProba)
	http.HandleFunc("/sleep", handlerSleep)

	magenta := color.New(color.FgHiMagenta).SprintFunc()
	fmt.Println(magenta("hello"))
	fmt.Println("Запускаю")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("!🤣🤣🤣🤣🤣🤣🤣🤣🤣🤣:", err.Error())
	}
	fmt.Println("Закончили")
}

// www.youtube.com/watch?v=4xST8IWJFZc
// 07.22
