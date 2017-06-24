package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/go-ini/ini"
)

// 生成32位MD5
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

//截取字符串
func Substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

//获取配置信息

func GetGlobalCfg() *ini.File {
	rootPath := GetRootDirectory()
	config, err := ini.Load(rootPath + "/.env")
	CheckError("fatal", err)
	return config
}
