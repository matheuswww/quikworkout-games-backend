package judge_util_service

import (
	"crypto/rand"
	"math/big"
)

func GenerateCode() (string, error) {
	max := big.NewInt(89999999)
	code, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}
	code.Add(code, big.NewInt(10000000))
	return code.String(), nil
}
