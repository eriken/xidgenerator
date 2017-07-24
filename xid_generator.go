package xidgenerator

import (
	"fmt"
	"math/rand"
)

const (
	X_ID_LENGHTH = 12
)

var (
	xIdChars = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func GenerateXiD() string {
	chars := make([]rune, X_ID_LENGHTH)
	for char := range chars {
		chars[char] = xIdChars[rand.Intn(len(xIdChars))]
	}
	return fmt.Sprintf("PUA-%s", string(chars))
}
