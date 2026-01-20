package simplecommection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func ChechConnection() {
	ctx := context.Background()
	con, err := pgx.Connect(ctx, "postgres://postgres:98098@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}
	if err := con.Ping(ctx); err != nil {
		panic(err)
	}
	fmt.Println("Успешное подключение к бд")
}
