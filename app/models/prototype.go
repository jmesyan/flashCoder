package models

type FlashBehavior struct {
	Bid       int64
	Bname     string
	Opid      int64
	Opname    string
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
	Tname   string
	Tcate   uint8
	Tsubs   string
	bids    string
	Addtime int64
	Updtime int64
}

type FlashTaskBehavior struct {
	Tbid     int64
	Bid      int64
	Bname    string
	Tid      int64
	Tname    string
	Ctid     int64
	CTname   string
	Border   int64
	Torder   int64
	Paramsin string
}

type FlashCron struct {
	Crid    int64
	Second  string
	Minute  string
	Hour    string
	Day     string
	Month   string
	Week    string
	Tid     int64
	Tname   string
	State   uint8
	Addtime int64
}
