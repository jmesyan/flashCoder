package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlModel struct {
	dbHandler *sql.DB
	connstr  string
}

func(m *MysqlModel)init(connstr string){
	var err error
	m.dbHandler,err=sql.Open("mysql",connstr)
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer m.dbHandler.Close()
}

func(m *MysqlModel)Select(table string,field []string,condition []Condition){

}


