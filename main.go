package main

import (
	"context"
	simplecommection "demo2/feature_postgres/simple_connection"
	simplesql "demo2/feature_postgres/simple_sql"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	conn, err := simplecommection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	if err := simplesql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	err = simplesql.InsertRow(ctx, conn, simplesql.TaskModel{
		Title:       "Покормить кота",
		Description: "Дать вискас",
		Completed:   false,
		CreatedAt:   time.Now(),
	})
	if err != nil {
		panic(err)
	}

	tasks, err := simplesql.SelectRow(ctx, conn)
	if err != nil {
		panic(err)
	}

	for _, task := range tasks {
		if task.Id == 1 {
			task.Title = "Покормить песика"
			task.Description = "Отсыпать корма"
			task.Completed = true
			now := time.Now()
			task.CompletedAt = &now

			if err := simplesql.UpdateRow(ctx, conn, task); err != nil {
				panic(err)
			}
			break
		}
	}
	fmt.Println("Succeed!")
}
