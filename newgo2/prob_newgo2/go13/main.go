package main

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

func handler(w http.ResponseWriter, r *http.Request) {

	str := "Hello, Eysk!"
	b := []byte(str)
	_, err := w.Write(b)

	if err != nil {
		fmt.Println("ERROR;", err.Error())

	} else {
		fmt.Println("😍😍😍-OK-😍😍")
	}
}

func cancelHandler(w http.ResponseWriter, r *http.Request) {

	str := "Pay exit!"
	b := []byte(str)
	_, err := w.Write(b)

	if err != nil {
		fmt.Println("ERROR;", err.Error())

	} else {
		fmt.Println("😂😂😂--OK--😂😂😂")
	}
}
func payHandler(w http.ResponseWriter, r *http.Request) {

	str := "New pay processed!"
	b := []byte(str)
	_, err := w.Write(b)

	if err != nil {
		fmt.Println("ERROR;", err.Error())

	} else {€
		fmt.Println("🎈🎈🎈--OK--🎈🎈🎈")
	}
}
func main() {
	red := color.New(color.FgHiRed).SprintFunc()
	fmt.Println(red("-----🚕🚓🚕-----"))

	http.HandleFunc("/default", handler)
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/cancel", cancelHandler)

	fmt.Println(red("Запускаю сервер🏦🏦🏦"))
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("ERRoooooOR:", err.Error())

	}
	fmt.Println(red("ЗАКОНЧИЛА ПРОГРАММА ВЫПОЛНЕНИЕ"))
}
