package flashdb

import (
	"database/sql"
	"encoding/json"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)

type FSqliDB struct {
	connstr   string
	dbHandler *sql.DB
}

func (m *FSqliDB) Init(connstr string) error {
	m.dbHandler, err = sql.Open("sqlite3", connstr)
	if err != nil {
		return err
	}
	m.connstr = connstr
	return nil
	return nil
}

func (m *FSqliDB) Close() {
	m.dbHandler.Close()
}

func (m *FSqliDB) Select(sql string, params []interface{}) (string, error) {
	if strings.Count(sql, "?") != len(params) {
		return "", errors.New("sql: params nums doesn't match need band")
	}
	stmtOut, err := m.dbHandler.Prepare(sql)
	if err != nil {
		return "", err
	}
	defer stmtOut.Close()
	rows, err := stmtOut.Query(params...)
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func (m *FSqliDB) SelectOne(sql string, params []interface{}, res []interface{}) error {
	if strings.Count(sql, "?") != len(params) {
		return errors.New("sql: params nums doesn't match need band")
	}
	stmtOut, _ := m.dbHandler.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmtOut.Close()
	return stmtOut.QueryRow(params...).Scan(res...)
}

func (m *FSqliDB) Insert(sql string, params []interface{}) (int64, error) {
	stmtIns, err := m.dbHandler.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmtIns.Close()
	res, err := stmtIns.Exec(params...)
	if err != nil {
		return 0, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastId, nil
}

func (m *FSqliDB) Update(sql string, params []interface{}) error {
	stmtIns, err := m.dbHandler.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	_, err = stmtIns.Exec(params...)
	if err != nil {
		return err
	}
	return nil
}

func (m *FSqliDB) TransBegin() (*sql.Tx, error) {
	tx, err := m.dbHandler.Begin()
	return tx, err
}

func (m *FSqliDB) TransInsert(tx *sql.Tx, sql string, params []interface{}) (int64, error) {
	stmtIns, err := tx.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmtIns.Close()
	res, err := stmtIns.Exec(params...)
	if err != nil {
		return 0, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastId, nil
}

func (m *FSqliDB) TransUpdate(tx *sql.Tx, sql string, params []interface{}) error {
	stmtIns, err := tx.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	_, err = stmtIns.Exec(params...)
	if err != nil {
		return err
	}
	return nil
}
