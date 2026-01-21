package simplesql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func SelectRow(ctx context.Context, conn *pgx.Conn) ([]TaskModel, error) {
	sqlRequest := `
SELECT id, title, description, completed, created_at, completed_at
FROM tasks
ORDER BY id ASC`

	rows, err := conn.Query(ctx, sqlRequest)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]TaskModel, 0)

	for rows.Next() {
		var task TaskModel

		err := rows.Scan(
			&task.Id,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
			&task.CompletedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
