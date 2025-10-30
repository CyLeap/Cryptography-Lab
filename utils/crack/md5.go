package crack

import (
	"crypto/md5"
	"fmt"
)

func CheckMD5(password string, targetHash string) bool {
	hash := md5.Sum([]byte(password))
	hashStr := fmt.Sprintf("%x", hash)
	return hashStr == targetHash
}
