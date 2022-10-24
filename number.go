package strgen

import "math/rand"

func Number(length int) string {
	const charset = "0123456789"
	code := make([]byte, length)
	i := 0
	for i < length {
		randomChar := charset[rand.Intn(len(charset))]
		code[i] = randomChar
		i++
	}

	return string(code)
}
