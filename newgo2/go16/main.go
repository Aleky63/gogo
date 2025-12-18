package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/fatih/color"
)

type Payment struct {
	Description string
	USD         int
}

var mtx = sync.Mutex{}
var money = 1000
var paymenHistory = make([]Payment, 0)

func payHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	httpRequestBodyString := string(httpRequestBody)
	parts := strings.SplitN(httpRequestBodyString, ",", 2)
	if len(parts) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func main() {

	red := color.New(color.FgHiRed).SprintFunc()
	fmt.Println(red("-----🚕🚓🚕-----"))

	http.HandleFunc("/pay", payHandler)

	// err := http.ListenAndServe(":1991", nil)
	// if err != nil {
	// 	fmt.Println("ERROR  during HTTP server operation:", err)
	// }

	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("ERROR  during HTTP server operation:", err)
	}

	fmt.Println("")
}
