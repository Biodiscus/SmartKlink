package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func HashMD5(in string) string {
	hash := md5.Sum([]byte(in))
	return hex.EncodeToString(hash[:])
}
