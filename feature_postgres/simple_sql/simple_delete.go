package simplesql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteRow(ctx context.Context, conn *pgx.Conn, tasksIDs []int) error {
	sqlRequest := `DELETE FROM tasks
WHERE id = ANY($1);
`
	_, err := conn.Exec(ctx, sqlRequest, tasksIDs)
	return err
}
