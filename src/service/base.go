package service

import (
	"context"
	"os"
	"synapsis-ecommerce/src/helper/validator"
	"synapsis-ecommerce/src/repository"

	"github.com/golang-jwt/jwt/v4"
)

type Service interface {
	getJWTClaims(cookie string) (*jwt.RegisteredClaims, error)
	Register(ctx context.Context, user validator.UserRegister) error
	Login(ctx context.Context, user validator.UserLogin) (string, error)
	GetProducts(ctx context.Context) ([]repository.Product, error)
}

type service struct {
	repo repository.Querier
}

func NewService(repo repository.Querier) Service {
	return &service{
		repo: repo,
	}
}

func (u *service) getJWTClaims(cookie string) (*jwt.RegisteredClaims, error) {

	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	return claims, nil

}
