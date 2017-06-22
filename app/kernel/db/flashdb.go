package flashdb

import (
	"database/sql"
	"flashCoder/utils"
)

var err error

type DriverType int

// Enum the Database driver
const (
	_          DriverType = iota // int enum type
	DRMySQL                      // mysql
	DRSqlite                     // sqlite
	DROracle                     // oracle
	DRPostgres                   // pgsql
	DRTiDB                       // TiDB
)

type FlashDB interface {
	Init(connstr string) error
	Insert(sql string, params []interface{}) (int64, error)
	Update(sql string, params []interface{}) error
	Select(sql string, params []interface{}) (string, error)
	SelectOne(sql string, params []interface{}, res []interface{}) error
	TransBegin() (*sql.Tx, error)
	TransInsert(tx *sql.Tx, sql string, params []interface{}) (int64, error)
	TransUpdate(tx *sql.Tx, sql string, params []interface{}) error
	Close()
}

func SetDbHandler(dbType DriverType, connstr string) FlashDB {
	switch dbType {
	case DRMySQL:
		db := new(FMyDB)
		err := db.Init(connstr)
		utils.CheckError(err)
		return db
	case DRSqlite:
		db := new(FSqliDB)
		err := db.Init(connstr)
		utils.CheckError(err)
		return db
	}
	return nil
}

func DBConnection(t string) (DriverType, string) {
	config := utils.GetGlobalCfg()
	dbType, dbconnstr := DRSqlite, ""
	dbconnect := config.Section("dbconnection").Key(t).String()
	switch dbconnect {
	case "sqlite":
		dbType = DRSqlite
		dbconnstr = config.Section("sqlite").Key("path").String()
	case "mysqld":
		dbType = DRMySQL
		dbmysql := config.Section("mysqld")
		dbconnstr = dbmysql.Key("user").String() + ":" + dbmysql.Key("passwd").String() + "@tcp(" + dbmysql.Key("host").String() + ":" + dbmysql.Key("port").String() + ")/" + dbmysql.Key("dbname").String()
	}
	return dbType, dbconnstr
}
