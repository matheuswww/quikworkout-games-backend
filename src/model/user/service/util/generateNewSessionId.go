package user_service_util

import "crypto/rand"

const idLength = 8
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateNewSessionId() (string, error) {
	b := make([]byte, idLength)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	for i := 0; i < idLength; i++ {
		b[i] = charset[int(b[i])%len(charset)]
	}
	return string(b), nil
}