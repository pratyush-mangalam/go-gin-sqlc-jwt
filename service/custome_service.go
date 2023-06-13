package service

import (
	"fmt"
	"go-boiler-plate/config"
	"go-boiler-plate/db"
	dbsql "go-boiler-plate/db/sqlc"
	"go-boiler-plate/jwt"
	"go-boiler-plate/model"
	"go-boiler-plate/util"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	Create(ctx *gin.Context, store db.Store, tokenRequest *model.UserSignUpRequest) (*model.UserSignInUpResponse, error)
}

type UserServiceImpl struct {
	Store db.Store
	Conf  *config.AppConfig
}

func NewUserServiceImpl(s db.Store, config *config.AppConfig) *UserServiceImpl {
	fs := &UserServiceImpl{Store: s, Conf: config}
	return fs
}

func (s *UserServiceImpl) SignUp(c *gin.Context, input *model.UserSignUpRequest) (*model.UserSignInUpResponse, error) {
	hashPassword, err := util.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}
	userDBParam := dbsql.CreateUserParams{Name: input.Name, Email: input.Email, Number: input.Number, Password: hashPassword}
	userData, err := s.Store.CreateUser(c, userDBParam)
	if err != nil {
		return nil, err
	}
	token, err := jwt.GenerateToken(int(userData.ID), s.Conf.JWTTokenLifeSpan, s.Conf.JWTTokenSecretKey)
	if err != nil {
		return nil, err
	}
	return &model.UserSignInUpResponse{Name: userData.Name, Email: userData.Email, Number: userData.Number, Token: token}, nil
}

func (s *UserServiceImpl) SignIn(c *gin.Context, input *model.UserSignInRequest) (*model.UserSignInUpResponse, error) {
	userData, err := s.Store.GetUserByEmail(c, input.Email)
	if err != nil {
		return nil, err
	}
	isEqual := util.CheckPasswordHash(input.Password, userData.Password)
	if !isEqual {
		return nil, fmt.Errorf("password doesn't match")
	}
	token, err := jwt.GenerateToken(int(userData.ID), s.Conf.JWTTokenLifeSpan, s.Conf.JWTTokenSecretKey)
	if err != nil {
		return nil, err
	}
	return &model.UserSignInUpResponse{Name: userData.Name, Email: userData.Email, Number: userData.Number, Token: token}, nil

}

func (s *UserServiceImpl) GetAllUser(c *gin.Context) (*model.GetAllUserResp, error) {
	usersData, err := s.Store.GetAllUsers(c)
	if err != nil {
		return nil, err
	}
	allUsersData := &model.GetAllUserResp{}
	for _, user := range usersData {
		allUsersData.Users = append(allUsersData.Users, &model.AllUsers{Name: user.Name, Email: user.Email, Number: user.Number})
	}
	return allUsersData, nil
}
