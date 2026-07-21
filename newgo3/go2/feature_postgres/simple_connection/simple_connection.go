package simple_connection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CheckConnection(ctx context.Context) (*pgx.Conn, error) {
	fmt.Println("The connection to the database was successful!")
	return pgx.Connect(ctx, "postgres://postgres:1803@localhost:5432/postgres")