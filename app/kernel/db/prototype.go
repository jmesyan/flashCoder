package flashdb

type FlashBehavior struct {
	Bid       int64
	Opid      int64
	Bname     string
	Paramsdef string
	Remark    string
	Addtime   int64
	Updtime   int64
}

type FlashOperate struct {
	Opid    int64
	Opname  string
	Optag   string
	Remark  string
	Addtime int64
}

type FlashTask struct {
	Tid     int64
	Tcate   uint8
	Tsubs   string
	bids    string
	Addtime int64
	Updtime int64
}

type FlashTaskBehavior struct {
	Tbid     int64
	Bid      int64
	Tid      int64
	Ctid     int64
	Border   int64
	Torder   int64
	Paramsin string
}

type FlashCron struct {
	Crid   int64
	Second uint8
	Minute uint8
	Hour   uint8
	Day    uint8
	Month  uint8
	Week   uint8
	Tid    int64
}
