package main

import (
	"study/feature_postgres/simple_connection"
)

func main() {
	ctx := contex.Background()
	conn, err := simple_connection.CheckConnection()
}
