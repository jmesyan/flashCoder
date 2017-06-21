package tester

import (
	"fmt"
	"github.com/go-ini/ini"
	"testing"
)

func TestIni(t *testing.T) {
	cfg, err := ini.Load("./../.env")
	fmt.Println(cfg, err)
	mysqld, err := cfg.GetSection("mysqld")
	fmt.Println(mysqld, err)
	host, err := mysqld.GetKey("host")
	fmt.Println(host, err)
}
