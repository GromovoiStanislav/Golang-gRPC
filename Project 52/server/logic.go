package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
)

// gravatarHash возвращает хэш-сумму MD5 для заданной строки.
func gravatarHash(email string) []byte {
	h := md5.New()
	io.WriteString(h, strings.ToLower(email))
	return h.Sum(nil)
}

// gravatarURL возвращает URL-адрес Gravatar на основе хэша и размера изображения.
func gravatarURL(hash []byte, size int32) string {
	return fmt.Sprintf("https://www.gravatar.com/avatar/%x?s=%d", hash, size)
}

// gravatar возвращает URL-адрес Gravatar для заданной электронной почты и размера изображения.
func gravatar(email string, size int32) string {
	hash := gravatarHash(email)
	return gravatarURL(hash, size)
}
