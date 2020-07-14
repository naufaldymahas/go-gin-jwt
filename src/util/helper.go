package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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

func TokenGenerator(email string) (string, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Load Env Failed:", err)
		return "", err
	}

	ISSUER := os.Getenv("ISSUER")

	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		Subject:   email,
		Issuer:    ISSUER,
	}

	SIGN_KEY := []byte(os.Getenv("SIGN_KEY"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(SIGN_KEY)
	if err != nil {
		log.Println("Signed String Failed:", err)
		return "", err
	}
	return ss, nil
}

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := godotenv.Load(); err != nil {
			log.Println("Unable load env")
			c.AbortWithStatus(http.StatusBadRequest)
		}
		auth := c.GetHeader("Authorization")

		if auth == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		} else if !strings.Contains(auth, "Bearer") {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ss := strings.Split(auth, " ")[1]

		SIGN_KEY := []byte(os.Getenv("SIGN_KEY"))

		token, err := jwt.Parse(ss, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return SIGN_KEY, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			log.Println(claims)
		} else {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
