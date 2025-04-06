package auth

import (
	"fmt"
	"golang_graphs/backend/internal/dto"
	"log"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

type Service interface {
	CreateToken(user dto.User) (string, error)
	ParseToken(tokenString string) (dto.User, error)
}

type service struct {
	secret []byte
}

func New(secret string) Service {
	return &service{secret: []byte(secret)}
}

func (s *service) CreateToken(user dto.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         strconv.FormatInt(user.Id, 10),
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"role":       user.Role,
	})

	return token.SignedString(s.secret)
}

func (s *service) ParseToken(tokenString string) (dto.User, error) {
	log.Printf("Token string %s", tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return s.secret, nil
	})
	if err != nil {
		return dto.User{}, errors.Wrap(err, "get token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		str, ok := claims["id"].(string)
		if !ok {
			return dto.User{}, fmt.Errorf("claim parsing id failed")
		}
		userID, err := strconv.ParseInt(str, 10, 64)
		log.Printf("userID %d", userID)
		if err != nil {
			return dto.User{}, fmt.Errorf("claim parsing failed")
		}
		firtsName, ok := claims["first_name"].(string)
		if !ok {
			return dto.User{}, fmt.Errorf("claim parsing first name failed")
		}
		lastName, ok := claims["last_name"].(string)
		if !ok {
			return dto.User{}, fmt.Errorf("claim parsing last name failed")
		}
		role, ok := claims["role"].(string)
		if !ok {
			return dto.User{}, fmt.Errorf("claim parsing role failed")
		}

		return dto.User{
			Id:        userID,
			FirstName: firtsName,
			LastName:  lastName,
			Role:      role,
		}, nil
	}

	return dto.User{}, fmt.Errorf("claim missing")
}
