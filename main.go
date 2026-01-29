package main

import (
	"context"
	simplecommection "demo2/feature_postgres/simple_connection"
	simplesql "demo2/feature_postgres/simple_sql"
	"fmt"
	"time"

	"github.com/k0kubun/pp"
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

	tasks, err := simplesql.SelectRow(ctx, conn)
	if err != nil {
		panic(err)
	}
	pp.Println(tasks)

	//err = simplesql.InsertRow(ctx, conn, simplesql.TaskModel{
	//	Title:       "Покормить выдру",
	//	Description: "Выдра голодна, очень хочет есть - так дай же",
	//	Completed:   false,
	//	CreatedAt:   time.Now(),
	//})
	//if err != nil {
	//	panic(err)
	//}

	simplesql.DeleteRow(ctx, conn, []int{5})

	for _, task := range tasks {
		if task.Id == 3 {
			task.Title = "Выиграть битву"
			task.Description = "Получить победу и кайф"
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
