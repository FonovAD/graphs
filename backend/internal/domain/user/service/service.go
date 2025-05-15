package userservice

import (
	"context"
	"fmt"
	model "golang_graphs/backend/internal/domain/model/user"
	"log"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

type UserService interface {
	CreateToken(user *model.User) (string, error)
	ParseToken(tokenString string) (*model.User, error)
}

type userService struct {
	secret []byte
}

func NewUserService(secret string) UserService {
	return &userService{secret: []byte(secret)}
}

func (us *userService) CreateToken(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":          strconv.FormatInt(user.Id, 10),
		"first_name":  user.FirstName,
		"last_name":   user.LastName,
		"father_name": user.FatherName,
		"role":        user.Role,
	})

	return token.SignedString(us.secret)
}

func (us *userService) ParseToken(tokenString string) (*model.User, error) {
	log.Printf("Token string %s", tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return us.secret, nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "get token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		str, ok := claims["id"].(string)
		if !ok {
			return nil, fmt.Errorf("claim parsing id failed")
		}
		userID, err := strconv.ParseInt(str, 10, 64)
		log.Printf("userID %d", userID)
		if err != nil {
			return nil, fmt.Errorf("claim parsing failed")
		}
		firtsName, ok := claims["first_name"].(string)
		if !ok {
			return nil, fmt.Errorf("claim parsing first name failed")
		}
		lastName, ok := claims["last_name"].(string)
		if !ok {
			return nil, fmt.Errorf("claim parsing last name failed")
		}
		fatherName, ok := claims["father_name"].(string)
		if !ok {
			return nil, fmt.Errorf("claim parsing father name failed")
		}
		role, ok := claims["role"].(string)
		if !ok {
			return nil, fmt.Errorf("claim parsing role failed")
		}

		return &model.User{
			Id:         userID,
			FirstName:  firtsName,
			LastName:   lastName,
			FatherName: fatherName,
			Role:       role,
		}, nil
	}

	return nil, fmt.Errorf("claim missing")
}

func (us *userService) AuthToken(ctx context.Context, token string) (*model.User, error) {
	user, err := us.ParseToken(token)
	if err != nil {
		return nil, err
	}

	return user, nil
}
