package jwt_usecase

import (
	"errors"
	"fmt"
	"fse-onboarding/repository/user_repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type JwtUsecaseImpl struct {
	UserRepository user_repository.UserRepository
}

func NewJwtUsecase(userRepository user_repository.UserRepository) JwtUsecase {
	return &JwtUsecaseImpl{
		UserRepository: userRepository,
	}
}

type CustomClaim struct {
	jwt.RegisteredClaims
	Role   string `json:"role"`
	UserID string `json:"user_id"`
}

func (jwtAuth *JwtUsecaseImpl) GenerateToken(userId string, roleId string) (string, error) {
	data, err := jwtAuth.UserRepository.GetRoleById(roleId)
	if err != nil {
		return "User not found", err
	}

	claim := CustomClaim{
		UserID: userId,
		Role:   data.Title,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 60)),
			Issuer:    os.Getenv("APP_NAME"),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claim)
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func (jwtAuth *JwtUsecaseImpl) ValidateToken(token string) (*jwt.Token, error) {

	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
}

func (jwtAuth *JwtUsecaseImpl) ValidateTokenAndGetUserId(token string) (string, error) {
	validatedToken, err := jwtAuth.ValidateToken(token)
	if err != nil {
		return "", err
	}

	claims, ok := validatedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("failed to claim token")
	}

	return claims["user_id"].(string), nil
}

func (jwtAuth *JwtUsecaseImpl) ValidateTokenAndGetRole(token string) (string, string, error) {
	validatedToken, err := jwtAuth.ValidateToken(token)
	if err != nil {
		return "", "", err
	}

	claims, ok := validatedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", errors.New("failed to claim token")
	}

	user, err := jwtAuth.UserRepository.GetDetailUsers(claims["user_id"].(string))

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return "", "", errors.New("user not found")
	}

	return user.Id, user.Role.Title, nil
}
