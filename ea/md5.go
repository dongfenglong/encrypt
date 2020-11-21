package ea

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(text string) string {
	m := md5.New()
	m.Write([]byte(text))
	return hex.EncodeToString(m.Sum(nil))
}
