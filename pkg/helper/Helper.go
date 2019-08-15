package helper

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	models "bitbucket.org/qasir-id/supplier-dashboard-service/database/models/user"
	jwt "github.com/dgrijalva/jwt-go"
)

var APPLICATION_NAME = "https://qasir.id"
var LOGIN_EXPIRATION_DURATION = time.Duration(6) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("cWFzaXIuaWQ6cHVibGlja2V5")

type TokenClaim struct {
	jwt.StandardClaims
	UserID     int64 `json:"user_id"`
	MerchantID int64 `json:"merchant_id"`
	OutletID   int64 `json:"outlet_id"`
}

func hashPassword(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	return string(hash)
}

func CheckHashedPassword(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	byePlainPwd := []byte(plainPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, byePlainPwd)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func GenerateJWT(user models.User) (string, error) {
	claims := TokenClaim{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		UserID:     int64(user.ID),
		MerchantID: user.MerchantID,
	}

	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)

	tokenString, err := token.SignedString(JWT_SIGNATURE_KEY)

	return tokenString, err
}

// func ValidateJWT(jwtToken string) (string, error) {
// 	claims := TokenClaim{}

// 	tkn, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
// 		return JWT_SIGNATURE_KEY, nil
// 	})

// 	if err != nil {

// 	}
// }
