package utils

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"iwexlmsapi/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/pbkdf2"
)

const iterations = 4096

func GenHash(password string) (string, string, error) {
	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		return "", "", err
	}
	saltStr := hex.EncodeToString(salt)
	hash := pbkdf2.Key([]byte(password), []byte(saltStr), iterations, sha512.Size, sha512.New)
	hashStr := hex.EncodeToString(hash)
	return hashStr, saltStr, nil
}

func ValidPassword(password string, dbHash string, dbSalt string) bool {
	hashVerify := pbkdf2.Key([]byte(password), []byte(dbSalt), iterations, sha512.Size, sha512.New)
	hashVerifyStr := hex.EncodeToString(hashVerify)
	return hashVerifyStr == dbHash
}

var PrivateKey *rsa.PrivateKey

func InitKeys() {
	privateKeyFile, err := os.Open("keys/id_rsa_priv.pem")
	if err != nil {
		panic(err.Error())
	}
	pemfileinfo, _ := privateKeyFile.Stat()
	size := pemfileinfo.Size()
	pembytes := make([]byte, size)
	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)
	data, _ := pem.Decode(pembytes)
	privateKeyFile.Close()
	PrivateKey, err = x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {
		panic(err.Error())
	}
}

type tokenStruct struct {
	Token     string `json:"token"`
	ExpiresIn string `json:"expiresIn"`
}

func IssueJWT(user *models.UserDB) (*tokenStruct, error) {
	token := jwt.New(jwt.GetSigningMethod("RS256"))
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.Email
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["authorized"] = true
	claims["role_name"] = user.RoleName
	claims["id"] = user.Id

	tokenString, err := token.SignedString(PrivateKey)
	if err != nil {
		return nil, err
	}
	return &tokenStruct{
		Token:     "Bearer " + tokenString,
		ExpiresIn: fmt.Sprintf("%dd", 1),
	}, nil
}
