package main

import (
	"context"

	"study/feature_postgres/simple_connection"
	"study/feature_postgres/simple_sql"
)

func main() {
	ctx := context.Background()
	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}
	simple_sql.CreateTable(ctx, conn)
}
