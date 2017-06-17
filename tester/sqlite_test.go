package tester

import (
	// "database/sql"
	"fmt"
	// _ "github.com/mattn/go-sqlite3"
	"log"
	// "os"
	"flashCoder/app/kernel/db"
	"testing"
)

type Users struct {
	UserId int
	Uname  string
}

func TestSqlite(t *testing.T) {
	// os.Remove("../foo.db")
	DB := flashdb.SetDbHandler(flashdb.DRSqlite, "../app/databases/flashCoder.db")

	sql := `insert into users(userId,uname) values(3,?);`
	params := []interface{}{"jamesyan"}
	DB.Insert(sql, params)
	// sql = `insert into users(userId,uname) values(2,'John');`
	// db.Exec(sql)
	rows, err := DB.Select("select * from users", []interface{}{})
	if err != nil {
		log.Fatal(err)
	}
	// defer rows.Close()
	// var users []Users = make([]Users, 0)
	// for rows.Next() {
	// 	var u Users
	// 	rows.Scan(&u.UserId, &u.Uname)
	// 	users = append(users, u)
	// }
	fmt.Println(rows)
}
