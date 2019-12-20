package util

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Str(src string) string {
	h := md5.New()
	h.Write([]byte(src)) // 需要加密的字符串为
	return hex.EncodeToString(h.Sum(nil))
}

func MD5Password(p string,s int) string {
	return MD5Str(MD5Str(p+string(s)))
}
