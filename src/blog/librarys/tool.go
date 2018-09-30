package librarys

import (
	"crypto/md5"
	"encoding/hex"
)

func StrToMd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	string := hex.EncodeToString(h.Sum(nil))
	return string
}

func PwdHash(s string) string {
	return StrToMd5(s)
}
