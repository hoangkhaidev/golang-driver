package repository

import (
	"context"
	"my-driver/model"
	"my-driver/model/req"
)

type UserRepo interface {
	SaveUser(context context.Context, user model.User) (model.User, error)
	CheckLogin(context context.Context, loginReq req.ReqSignIn) (model.User, error)
	SelectUserById(context context.Context, userId string) (model.UserResponse, error)
	UpdateUser(context context.Context, user model.UserResponse) (model.UserResponse, error)
}