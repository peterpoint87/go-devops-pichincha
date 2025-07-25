package jwtutil

import (
    "os"
    "time"
    "log"
    "github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte

func init() {
    secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        log.Fatal("Environment variable JWT_SECRET not set")
    }
    jwtSecret = []byte(secret)
}

func GenerateJWT(to string, ttlSec int) string {
    claims := jwt.MapClaims{
        "to":  to,
        "exp": time.Now().Add(time.Duration(ttlSec) * time.Second).Unix(),
        "iat": time.Now().Unix(),
        "kid": "devops-v1",
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signed, _ := token.SignedString(jwtSecret)
    return signed
}