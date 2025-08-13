package main

import "fmt"

type Event struct {
	Title string
	Date string
    Location string
}


func createGoEvent() Event{
    event := Event {
	Title: "День рождения Golang",
   Date: "10 ноября 2009",
    Location: "GoogleLand",
}
	return event
}


func main46() {
	
 res:= createGoEvent()

	fmt.Println(res)
	fmt.Printf("Событие: %s\nДата: %s\nМесто: %s\n", res.Title, res.Date, res.Location)
}

// https://stepik.org/lesson/1500866/step/4?unit=1520983