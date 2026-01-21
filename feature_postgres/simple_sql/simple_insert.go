package simplesql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, conn *pgx.Conn, task TaskModel) error {
	sqlRequest := `INSERT INTO tasks (title, description, completed, created_at, completed_at)
VALUES($1, $2, $3, $4, $5);`

	// 5 параметров в VALUES, значит 5 аргументов
	_, err := conn.Exec(ctx, sqlRequest,
		task.Title,
		task.Description,
		task.Completed,
		task.CreatedAt,
		task.CompletedAt)
	return err
}
