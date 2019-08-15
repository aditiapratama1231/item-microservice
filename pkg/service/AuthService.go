package service

import (
	"context"

	"github.com/jinzhu/gorm"

	models "bitbucket.org/qasir-id/supplier-dashboard-service/database/models/user"
	"bitbucket.org/qasir-id/supplier-user-service/pkg/helper"
	payload "bitbucket.org/qasir-id/supplier-user-service/pkg/request/payload"
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

func (userService) LoginUser(ctx context.Context, data payload.LoginRequest) (payload.LoginResponse, error) {
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
