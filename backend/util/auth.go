package util

import (
	"golang.org/x/crypto/bcrypt"
)

// GeneratePasswordHash 產生密碼 Hash 加密
func GeneratePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
