package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func HashPassword(password string) string {
	if err := godotenv.Load(); err != nil {
		log.Panic(err.Error())
	}
	s := os.Getenv("SECRET_KEY")

	h := hmac.New(sha256.New, []byte(s))
	h.Write([]byte(password))

	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}
