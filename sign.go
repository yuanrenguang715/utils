package utils

import (
	"strings"
)

const (
	DesKey = "r5k1*8a$@8dc!dytkcs2dqz!"
)

func VerifySign(body, key string) string {
	return strings.ToLower(MD5(strings.ToLower(MD5(body)) + key))
}
