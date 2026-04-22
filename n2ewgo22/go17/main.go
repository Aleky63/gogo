package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.URL.Path)
}

func main() {
http.HandleFunc("/tasks/",handler)

if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("ERROR :", err)
	}
}