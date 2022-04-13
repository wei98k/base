package core

import (
	"crypto/md5"
	"io"
)

func toShortUrl(url string) string {
	h := md5.New()
	io.WriteString(h, url)
	m1 := h.Sum(nil)
	return string(m1)
}
