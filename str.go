package test

import (
	"math/rand"
	"time"
)

func GetRandomString(l int) string {
	str := "23456789abcdefghijkmnpqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
