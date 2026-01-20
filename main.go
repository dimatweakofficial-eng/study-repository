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
	if err := simplesql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}
	if err := simplesql.InsertRow(ctx, conn, "Повеселиться с крокодилом", "Хочу сильных эмоций бля, дайте мне его", false, time.Now()); err != nil {
		panic(err)
	}
	if err := simplesql.UpdateRow(ctx, conn); err != nil {
		panic(err)
	}
	if err := simplesql.DeleteRow(ctx, conn); err != nil {
		panic(err)
	}

	fmt.Println("Succesed!")
}
