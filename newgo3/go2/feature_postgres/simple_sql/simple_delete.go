package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeletetRow(ctx context.Context,
	conn *pgx.Conn,
) error {
	sqlQuery := `
DELETE FROM  tasks
	WHERE id=9;
	
	
	`
	_, err := conn.Exec(ctx, sqlQuery)

	return err
}
