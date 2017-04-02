package flashdb

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"errors"
)

var err error

type FMyDB struct {
	dbHandler *sql.DB
	connstr   string
}

func (m *FMyDB)Init(connstr string) {
	m.dbHandler, err = sql.Open("mysql", connstr)
	if err != nil {
		panic(err.Error())
	}
}

func (m *FMyDB)Close(){
	m.dbHandler.Close()
}

func (m *FMyDB)Select(sql string, params ...interface{}) (*sql.Rows, error) {
	if strings.Count(sql, "?") != len(params) {
		return nil,errors.New("sql: params nums doesn't match need band")
	}
	stmtOut, err := m.dbHandler.Prepare(sql)
	if err != nil {
		return nil,err
	}
	defer stmtOut.Close()
	return stmtOut.Query(params...)
}

func (m *FMyDB)SelectOne(sql string, params []interface{},res []interface{}) error {
	if strings.Count(sql, "?") != len(params) {
		return errors.New("sql: params nums doesn't match need band")
	}
	stmtOut, _ := m.dbHandler.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmtOut.Close()
	return stmtOut.QueryRow(params...).Scan(res ...)
}


