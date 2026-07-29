package main

import (
	"context"
	"fmt"
	"time"

	"study/feature_postgres/simple_connection"
	"study/feature_postgres/simple_sql"

	"github.com/fatih/color"
	"github.com/k0kubun/pp"
)

func main() {
	ctx := context.Background()
	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}
	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	// if err := simple_sql.InsertRow(
	// 	ctx,
	// 	conn,
	// 	"Lunch-new",
	// 	"Need to eat",
	// 	false,
	// 	time.Now(),
	// ); err != nil {
	// 	panic(err)
	// }

	// if err := simple_sql.UpdateRow(ctx, conn); err != nil {
	// 	panic(err)
	// }

	// if err := simple_sql.DeletetRow(ctx, conn); err != nil {
	// 	panic(err)
	// }
	tasks, err := simple_sql.SelectRows(ctx, conn)
	if err != nil {
		panic(err)
	}

	for _, task := range tasks {
		if task.ID == 7 {

			task.Title = " TRAMP😊😊😊😊😊"
			task.Description = "The red-haired horseradish is weird"
			task.Completed = true
			now := time.Now()
			task.CompletedAt = &now

			if err := simple_sql.UpdateTask(ctx, conn, task); err != nil {
				panic(err)
			}
			break
		}
	}

	red := color.New(color.FgRed).SprintFunc()
	fmt.Println(red("-----------------------"))
	fmt.Println(red("🛠️🛠️- Succeed!-🛠️🛠️"))
	pp.Println(tasks)
}
