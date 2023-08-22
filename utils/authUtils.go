package utils

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"iwexlmsapi/models"

	// "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/pbkdf2"
)

const iterations = 4096

func GenHash(password string) (string, string, error) {
	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		return "", "", err
	}
	hash := pbkdf2.Key([]byte(password), salt, iterations, sha512.Size, sha512.New)
	hashStr := hex.EncodeToString(hash)
	saltStr := hex.EncodeToString(salt)
	return hashStr, saltStr, nil
}

func ValidPassword(password string, dbHash string, dbSalt string) bool {
	hashVerify := pbkdf2.Key([]byte(password), []byte(dbSalt), iterations, sha512.Size, sha512.New)
	hashVerifyStr := hex.EncodeToString(hashVerify)
	return hashVerifyStr == dbHash
}

func IssueJWT(user *models.User) {

}
