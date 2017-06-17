package utils

import (
	"crypto/md5"
	"encoding/hex"
	"log"
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
