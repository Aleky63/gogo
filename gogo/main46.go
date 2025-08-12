package main

import "fmt"

type Event struct {
	Title string
	Date string
    Location string
}


func createGoEvent() Event{
return Event {

Title: "День рождения Golang",
Date: "10 ноября 2009",
Location: "GoogleLand",
   }
}


func main46() {
	
 res:= createGoEvent()

	fmt.Println(res)
	fmt.Printf("Событие: %s\nДата: %s\nМесто: %s\n", res.Title, res.Date, res.Location)
}

// https://stepik.org/lesson/1500866/step/4?thread=solutions&unit=1520983