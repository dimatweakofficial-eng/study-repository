package simplesql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteRow(ctx context.Context, conn *pgx.Conn) error {
	sqlQuerry := `DELETE FROM tasks
	WHERE id = 8;
	`
	_, err := conn.Exec(ctx, sqlQuerry)
	return err

}
