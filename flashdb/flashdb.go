package flashdb


type FlashDB interface {
	Init(connstr string)
	Close()
	Select(sql string ,params ...interface{})
	SelectOne(sql string ,params ...interface{})
}
