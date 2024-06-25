package core

import (
	"crypto/rand"
	"math/big"
	mathrandom "math/rand"
	"time"
)

var (
	character = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	chLen     = len(character)
)

func Random(size int) []byte {
	return []byte(RandomString(size))
}

func RandomString(size int) string {
	buf := make([]byte, size)
	maxB := big.NewInt(int64(chLen))
	for i := 0; i < size; i++ {
		random, err := rand.Int(rand.Reader, maxB)
		if err != nil {
			mathrandom.New(mathrandom.NewSource(time.Now().UnixNano()))
			buf[i] = character[mathrandom.Intn(chLen)]
			continue
		}
		buf[i] = character[random.Int64()]
	}
	return string(buf)
}
