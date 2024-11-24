package service

import (
	"context"
	"os"
	"strings"
	"synapsis-ecommerce/src/helper/validator"
	"synapsis-ecommerce/src/repository"

	"github.com/golang-jwt/jwt/v4"
)

type Service interface {
	getJWTClaims(cookie string) (*jwt.RegisteredClaims, error)
	Register(ctx context.Context, user validator.UserRegister) error
	Login(ctx context.Context, user validator.UserLogin) (string, error)
	GetProducts(ctx context.Context) ([]repository.Product, error)
	GetProductByCategory(ctx context.Context, category string) ([]repository.Product, error)
	GetCartByUserId(ctx context.Context, authHeader string) ([]repository.GetCartByUserIdRow, error)
	AddProductToCartTx(ctx context.Context, authHeader string, arg validator.AddProductCartUser) error
	DeleteProductOnCartById(ctx context.Context, authHeader string, cartItemId validator.DeleteProductOnCart) error
}

type service struct {
	repo repository.Store
}

func NewService(repo repository.Store) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) getJWTClaims(authHeader string) (*jwt.RegisteredClaims, error) {

	parts := strings.Split(authHeader, " ")

	tokenString := parts[1]

	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	return token.Claims.(*jwt.RegisteredClaims), nil

}
