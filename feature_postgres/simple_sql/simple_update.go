package simplesql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateRow(ctx context.Context, conn *pgx.Conn) error {
	sqlQuerry := `UPDATE tasks
SET completed = 'true'
WHERE id = 3
;`
	_, err := conn.Exec(ctx, sqlQuerry)
	return err
}
