package service

import (
	"context"
	"os"
	"strings"
	"synapsis-ecommerce/src/repository"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (s *service) GetCartByUserId(ctx context.Context, authHeader string) ([]repository.GetCartByUserIdRow, error) {

	parts := strings.Split(authHeader, " ")

	tokenString := parts[1]

	claims := &jwt.RegisteredClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	userID, err := uuid.Parse(claims.Issuer)
	if err != nil {
		return nil, err
	}

	cartUser, err := s.repo.GetCartByUserId(ctx, uuid.NullUUID{UUID: userID, Valid: true})
	if err != nil {
		return nil, err
	}

	return cartUser, nil

}
