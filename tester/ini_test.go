package tester

import (
	"fmt"
	"github.com/go-ini/ini"
	"testing"
)

func TestIni(t *testing.T) {
	cfg, err := ini.Load("./../.env")
	fmt.Println(cfg, err)
	mysqld, err := cfg.GetSection("staticDir")
	fmt.Println(mysqld, err)
	yes := mysqld.HasKey("/assets")
	fmt.Println(yes)
	host, err := mysqld.GetKey("/assets")
	fmt.Println(host, err)
}
