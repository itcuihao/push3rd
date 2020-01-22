package utils

import (
	"math/rand"
	"time"
)

const digitBytes = "0123456789"

//生成随机字符串
func RandDigitString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = digitBytes[rand.Intn(len(digitBytes))]
	}
	return string(b)
}

func GetRandomSalt() string {
	return RandDigitString(4)
}

var src = rand.NewSource(time.Now().UnixNano())

const alnumBytes = "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// RandStringBytes 生成字母随机数
func RandStringBytes(n int, candidate string) []byte {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(candidate) {
			b[i] = candidate[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return b
}

func GetRandStr(len int) string {
	return string(RandStringBytes(len, alnumBytes))
}
