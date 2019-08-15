package service

import (
	"context"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	models "bitbucket.org/qasir-id/supplier-dashboard-service/database/models/user"
	payload "bitbucket.org/qasir-id/supplier-user-service/pkg/request/payload"
	jwt "github.com/dgrijalva/jwt-go"
)

type UserService interface {
	LoginUser(context.Context, payload.LoginRequest) (payload.LoginResponse, error)
}

var query *gorm.DB

type userService struct{}

func NewUserService(db *gorm.DB) UserService {
	query = db
	return userService{}
}

var APPLICATION_NAME = "qasir.id"
var LOGIN_EXPIRATION_DURATION = time.Duration(6) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("cWFzaXIuaWQ6cHVibGlja2V5")

type TokenClaim struct {
	jwt.StandardClaims
	UserID     int64 `json:"user_id"`
	MerchantID int64 `json:"merchant_id"`
	OutletID   int64 `json:"outlet_id"`
}

func (userService) LoginUser(ctx context.Context, data payload.LoginRequest) (payload.LoginResponse, error) {
	var user models.User

	if query.Where("username = ?", data.Data.Username).Find(&user).RecordNotFound() {
		return payload.LoginResponse{
			Message:    "User Not found",
			StatusCode: 404,
			Err:        true,
		}, nil
	}

	if !checkHashedPassword(user.Password, data.Data.Password) {
		return payload.LoginResponse{
			Message:    "Failed to Login : Incorrect Password",
			StatusCode: 401,
			Err:        true,
		}, nil
	}

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

	if err != nil {
		return payload.LoginResponse{
			Message:    "Failed To login : Internal Server Error",
			StatusCode: 500,
			Err:        true,
		}, nil
	}

	return payload.LoginResponse{
		Message:    "Login Successfully",
		StatusCode: 200,
		Err:        true,
		Data: payload.LoginToken{
			AccessToken: tokenString,
			User:        user,
		},
	}, nil
}

func hashPassword(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	return string(hash)
}

func checkHashedPassword(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	byePlainPwd := []byte(plainPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, byePlainPwd)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
