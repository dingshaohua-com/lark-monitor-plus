package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Teacher struct {
	id   string
	name string
	age  int
}

func main() {
	var databaseUrl = "postgres://postgres:dshvv@103.110.80.247:5432/school"
	conn, _ := pgx.Connect(context.Background(), databaseUrl)
	defer conn.Close(context.Background())

	var user Teacher
	var sql = "select * from teacher where id=$1"
	conn.QueryRow(context.Background(), sql, 1).Scan(&user.id, &user.name, &user.age)
	fmt.Printf("%+v\n", user) // + 号表示 "Plus keys"（带上字段名）
}

//conn, _ := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
