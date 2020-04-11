package log

import (
	"crypto/rand"
	"encoding/hex"
)

// 缺省的id生成器
func IdGenerator() string {
	data := make([]byte, IdSize)
	if _, err := rand.Read(data); err != nil {
		return ""
	}

	return hex.EncodeToString(data)
}
