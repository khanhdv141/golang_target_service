package service

import (
	"CMS/dto"
	"CMS/repository"
	"CMS/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService interface {
	Login(*gin.Context, *dto.LoginRequest) *dto.BaseResponse[*dto.Token]
}

type authService struct {
	userRepo repository.UserRepository
	jwtUtils util.JWTUtils
}

func NewAuthService(userRepo repository.UserRepository, jwtUtils util.JWTUtils) AuthService {
	return &authService{
		userRepo: userRepo,
		jwtUtils: jwtUtils,
	}
}

func (s *authService) Login(ctx *gin.Context, req *dto.LoginRequest) *dto.BaseResponse[*dto.Token] {
	user, err := s.userRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		return MakeBadRequestResponse[*dto.Token]("User not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return MakeBadRequestResponse[*dto.Token]("Wrong password")
	}

	now := time.Now()
	claims := dto.JwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(30 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
		Username: user.Username,
		UserId:   user.ID,
	}
	token, err := s.jwtUtils.GenerateToken(claims)
	if err != nil {
		return MakeBadRequestResponse[*dto.Token]("Error when create token")
	}

	return MakeSuccessResponse[*dto.Token](&dto.Token{
		Token:     token,
		TokenType: "Bearer",
	})
}
