package service

import (
	"context"
	"strings"

	"github.com/jinzhu/gorm"

	models "bitbucket.org/qasir-id/supplier-dashboard-service/database/models/user"
	"bitbucket.org/qasir-id/supplier-user-service/pkg/helper"
	payload "bitbucket.org/qasir-id/supplier-user-service/pkg/request/payload"
	jwt "github.com/dgrijalva/jwt-go"
)

type AuthService interface {
	LoginUser(context.Context, payload.LoginRequest) (payload.LoginResponse, error)
	InstropectionToken(context.Context, payload.TokenInstropectionRequest) (payload.TokenInstropectionResponse, error)
}

var query *gorm.DB

type authService struct{}

func NewAuthService(db *gorm.DB) AuthService {
	query = db
	return authService{}
}

func (authService) LoginUser(ctx context.Context, data payload.LoginRequest) (payload.LoginResponse, error) {
	var user models.User

	if query.Where("username = ?", data.Data.Username).Find(&user).RecordNotFound() {
		return payload.LoginResponse{
			Message:    "User Not found",
			StatusCode: 404,
			Err:        true,
		}, nil
	}

	if !helper.CheckHashedPassword(user.Password, data.Data.Password) {
		return payload.LoginResponse{
			Message:    "Failed to Login : Incorrect Password",
			StatusCode: 401,
			Err:        true,
		}, nil
	}

	tokenString, err := helper.GenerateJWT(user)

	if err != nil {
		return payload.LoginResponse{
			Message:    "Failed To login : Internal Server Error",
			StatusCode: 500,
			Err:        true,
		}, nil
	}

	userData := payload.User{
		ID:         int64(user.ID),
		MerchantID: user.MerchantID,
		OutletID:   user.OutletID,
		Username:   user.Username,
		Email:      user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Title:      user.Title,
		Image:      user.Image,
		Level:      user.Level,
		Status:     user.Status,
	}

	return payload.LoginResponse{
		Message:    "Login Successfully",
		StatusCode: 200,
		Err:        true,
		Data: payload.LoginToken{
			AccessToken: tokenString,
			User:        userData,
		},
	}, nil
}

func (authService) InstropectionToken(ctx context.Context, data payload.TokenInstropectionRequest) (payload.TokenInstropectionResponse, error) {
	token := data.Token
	tokenString := strings.Replace(token, "Bearer ", "", -1)

	if token == "" {
		return payload.TokenInstropectionResponse{
			Error:      "invalid_request",
			Message:    "Missing Authorization Header",
			StatusCode: 400,
		}, nil
	}

	claims := &helper.TokenClaim{}

	tkn, _ := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return helper.JWT_SIGNATURE_KEY, nil
	})

	if !tkn.Valid || !strings.Contains(token, "Bearer") {
		return payload.TokenInstropectionResponse{
			Error:      "invalid_request",
			Message:    "Token Invalid",
			StatusCode: 403,
		}, nil
	}

	return payload.TokenInstropectionResponse{
		Activate: true,
		Issuer:   claims.Issuer,
		Exp:      claims.ExpiresAt,
		UserID:   claims.UserID,
		Sub:      claims.Subject,
	}, nil
}
