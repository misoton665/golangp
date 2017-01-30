package main

import (
	"database/sql"
	"fmt"
	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq"
)

type User struct {
	Id   int32
	Name string
}

func renameColumns(t *gorp.TableMap) {
	t.ColMap("Id").Rename("id")
	t.ColMap("Name").Rename("name")
}

func main() {
	dsn := "dbname=testdb host=localhost port=5432 user=misoton password=testpass"
	driver := "postgres"

	db, err := sql.Open(driver, dsn)

	if err == nil {
		fmt.Println("no error.")
	} else {
		panic("Error connecting to db: " + err.Error())
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	t := dbmap.AddTableWithName(User{}, "test_users").SetKeys(true, "Id")
	renameColumns(t)

	dbmap.DropTables()
	err = dbmap.CreateTables()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("hoghoge")
}
