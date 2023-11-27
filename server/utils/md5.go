package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// Md5Encode 转小写
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

// MD5Encode 转大写
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// MakePassword 加密
func MakePassword(password, salt string) string {
	return Md5Encode(password + salt)
}

// ValidPassword 解密
func ValidPassword(password, salt, sqlPassword string) bool {
	return Md5Encode(password+salt) == sqlPassword
}
