package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/naoina/genmai"
	"time"
)

type (
	// 旧テーブル（正規化前）
	CustomerPurchase struct {
		Id        int64 `db:"pk"`
		Name      string
		ItemId    int64
		ItemName  string
		Price     int64
		CreatedAt time.Time
	}

	// 新テーブル（正規化後）
	Customer struct {
		Id   int64  `db:"pk"`
		Name string `db:"unique"`
	}

	Item struct {
		Id    int64  `db:"pk"`
		Name  string `db:"unique"`
		Price int64
	}

	Purchase struct {
		Id         int64 `db:"pk"`
		CustomerId int64
		ItemId     int64
		CreatedAt  time.Time
	}
)

func panicIfError(maybeError error) {
	if maybeError != nil {
		panic(maybeError)
	}
}

func printIfError(maybeError error) {
	if maybeError != nil {
		fmt.Println(maybeError)
	}
}

func insertTestData(db *genmai.DB) {
	panicIfError(db.CreateTableIfNotExists(CustomerPurchase{}))

	rows := []CustomerPurchase{
		{Name: "user1", ItemId: 1, ItemName: "item1", Price: 100, CreatedAt: time.Now()},
		{Name: "user2", ItemId: 2, ItemName: "item2", Price: 200, CreatedAt: time.Now()},
		{Name: "user3", ItemId: 3, ItemName: "item3", Price: 300, CreatedAt: time.Now()},
		{Name: "user4", ItemId: 4, ItemName: "item4", Price: 400, CreatedAt: time.Now()},
		{Name: "user5", ItemId: 1, ItemName: "item1", Price: 100, CreatedAt: time.Now()},
		{Name: "user6", ItemId: 1, ItemName: "item1", Price: 100, CreatedAt: time.Now()},
		{Name: "user1", ItemId: 2, ItemName: "item2", Price: 200, CreatedAt: time.Now()},
		{Name: "user2", ItemId: 3, ItemName: "item3", Price: 300, CreatedAt: time.Now()},
		{Name: "user3", ItemId: 1, ItemName: "item1", Price: 100, CreatedAt: time.Now()},
		{Name: "user1", ItemId: 4, ItemName: "item4", Price: 400, CreatedAt: time.Now()},
		{Name: "user7", ItemId: 3, ItemName: "item3", Price: 300, CreatedAt: time.Now()},
		{Name: "user8", ItemId: 2, ItemName: "item2", Price: 200, CreatedAt: time.Now()},
		{Name: "user9", ItemId: 1, ItemName: "item1", Price: 100, CreatedAt: time.Now()},
		{Name: "user4", ItemId: 1, ItemName: "item1", Price: 100, CreatedAt: time.Now()},
	}

	_, err := db.Insert(rows)
	panicIfError(err)
}

func doNormalization(db *genmai.DB) {
	// 必要なテーブルの作成
	panicIfError(db.CreateTableIfNotExists(Customer{}))
	panicIfError(db.CreateTableIfNotExists(Item{}))
	panicIfError(db.CreateTableIfNotExists(Purchase{}))

	//var n int64
	//panicIfError(db.Select(&n, db.Count(), db.From(CustomerPurchase{})))
	//fmt.Println(n)

	// 旧テーブルのデータを取得
	var cps []CustomerPurchase
	panicIfError(db.Select(&cps))

	// 一つひとつ新テーブルへ格納
	for _, cp := range cps {
		// 新テーブル用のデータ作成
		customer := &Customer{Name: cp.Name}
		item := &Item{Id: cp.ItemId, Name: cp.ItemName, Price: cp.Price}

		// INSERT
		_, err1 := db.Insert(customer)
		printIfError(err1)
		_, err2 := db.Insert(item)
		printIfError(err2)

		var c []Customer

		panicIfError(db.Select(&c, db.Where("name", "=", cp.Name)))

		purchase := &Purchase{Id: cp.Id, CustomerId: c[0].Id, ItemId: cp.ItemId, CreatedAt: cp.CreatedAt}
		_, err3 := db.Insert(purchase)
		printIfError(err3)
	}

	db.DropTable(CustomerPurchase{})
}

func main() {
	db, err := genmai.New(&genmai.MySQLDialect{}, "user:pass@tcp(127.0.0.1:3306)/test?parseTime=true")

	panicIfError(err)

	defer db.Close()

	insertTestData(db)

	doNormalization(db)
}
