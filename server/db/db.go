package main

import (
	//	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/naoina/genmai"
	//	"time"
)

func main() {
	dsn := "dbname=testdb host=localhost port=5432 user=misoton password=testpass"

	db, err := genmai.New(&genmai.PostgresDialect{}, dsn)

	if err == nil {
		fmt.Println("no error.")
	} else {
		panic("Error connecting to db: " + err.Error())
	}
	defer db.Close()

	if err := db.CreateTable(&IbukiUser{}); err != nil {
		panic(err)
	}

	fmt.Println("hoghoge")
}
