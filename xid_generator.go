package xidgenerator

import (
	"fmt"
	"math/rand"
)

const (
	x_id_length = 12
)

var (
	xIdChars = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func GenerateXiD() string {
	chars := make([]rune, x_id_length)
	for char := range chars {
		chars[char] = xIdChars[rand.Intn(len(xIdChars))]
	}
	return fmt.Sprintf("PUA-%s", string(chars))
}
