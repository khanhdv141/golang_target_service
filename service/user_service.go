package service

import (
	"CMS/dto"
	"CMS/model"
	"CMS/repository"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(*gin.Context, *dto.CreateUserRequest) *dto.BaseResponse[*model.User]
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) CreateUser(ctx *gin.Context, req *dto.CreateUserRequest) *dto.BaseResponse[*model.User] {
	if req.Username == "" || req.Password == "" {
		return MakeBadRequestResponse[*model.User]("Username or password is empty")
	}
	dbUser, err := s.userRepository.FindByUsername(ctx, req.Username)
	if dbUser != nil {
		return MakeBadRequestResponse[*model.User]("Username already exists")
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	req.Password = string(bytes)
	var user = model.User{
		Username: req.Username,
		Password: req.Password,
	}
	err = s.userRepository.Save(ctx, &user)
	if err != nil {
		return MakeBadRequestResponse[*model.User]("Cannot save user")
	}
	return MakeSuccessResponse[*model.User](&user)
}

func (s *userService) DeleteUser(ctx *gin.Context, id uint) *dto.BaseResponse[*model.User] {
	user, err := s.userRepository.FindById(ctx, id)
	if user == nil {
		return MakeBadRequestResponse[*model.User]("User not found")
	}
	err = s.userRepository.DeleteById(ctx, id)
	if err != nil {
		return MakeBadRequestResponse[*model.User]("Cannot delete user")
	}
	return MakeSuccessResponse[*model.User](user)
}
