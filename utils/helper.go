package utils

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	// "os"
	// "math/rand"
	// "time"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

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
