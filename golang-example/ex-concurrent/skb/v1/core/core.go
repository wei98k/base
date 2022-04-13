package core

import (
	"crypto/md5"
	"fmt"
	"io"
)

// 长连接转换成短链接函数
func ToShortUrl(longUrl string) string {
	h := md5.New()
	io.WriteString(h, longUrl)
	return fmt.Sprintf("%x", string(h.Sum(nil)))
}
