package xidgenerator

import (
	"fmt"
	"math/rand"
)

const (
	defaultLength = 12
	defaultPrefix = "xid-"
)

var (
	length   int
	prefix   string
	allChars = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func Init(lengthParam int, prefixParam string) {
	length = lengthParam
	prefix = prefixParam
}

func init() {
	length = defaultLength
	prefix = defaultPrefix
}

func Generate() string {
	chars := make([]rune, length)
	for char := range chars {
		chars[char] = allChars[rand.Intn(len(allChars))]
	}
	return fmt.Sprintf("%s%s", prefix, string(chars))
}
