package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	INSERT INTO tasks (
		title ,
	description ,
	completed ,
	created_at )
	 VALUES ('Homework2','Do the task before the 10th number','false','2026-11-05 18:00:06'
		)`

	_, err := conn.Exec(ctx, sqlQuery)

	return err
}
