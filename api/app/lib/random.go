package lib

import "crypto/rand"

func CreateRandomStr(digit uint32) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, digit)

	rand.Read(b)

	var result string
	for _, v := range b {
		result += string(letters[int(v)%len(letters)])
	}
	return result
}
