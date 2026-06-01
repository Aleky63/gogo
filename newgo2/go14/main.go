package main

import (
	"fmt"
	"net/http"

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
		fmt.Println("OK_____________________OK")
	}
}

func main() {
	http.HandleFunc("/default", handler)
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/cancel", cancelHandler)

	magenta := color.New(color.FgHiMagenta).SprintFunc()
	fmt.Println(magenta("hello"))
	fmt.Println("Запускаю")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("!🤣🤣🤣🤣🤣🤣🤣🤣🤣🤣:", err.Error())
	}
	fmt.Println("Закончили")
}
