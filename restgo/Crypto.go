package restgo

import (

	"crypto/md5"
	"encoding/hex"
	"strings"
)



func Md5encode(arg0 string) string{
	h := md5.New()
	h.Write([]byte(arg0))
	cipherStr := h.Sum(nil)
	return strings.ToLower(hex.EncodeToString(cipherStr)) // 输出加密结果
}


