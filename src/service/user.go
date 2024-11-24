package service

import (
	"context"
	"synapsis-ecommerce/src/helper/validator"
	"synapsis-ecommerce/src/repository"

	"github.com/google/uuid"
)

func (s *service) GetCartByUserId(ctx context.Context, authHeader string) ([]repository.GetCartByUserIdRow, error) {

	claims, err := s.getJWTClaims(authHeader)
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

func (s *service) AddProductToCartTx(ctx context.Context, authHeader string, arg validator.AddProductCartUser) error {

	claims, err := s.getJWTClaims(authHeader)
	if err != nil {
		return err
	}

	uuidUser, err := uuid.Parse(claims.Issuer)
	if err != nil {
		return err
	}

	uuidProduct, err := uuid.Parse(arg.ProductId)
	if err != nil {
		return err
	}

	argRepo := repository.AddProductToCartParams{
		UserID: uuid.NullUUID{
			UUID:  uuidUser,
			Valid: true,
		},
		ProductID: uuid.NullUUID{
			UUID:  uuidProduct,
			Valid: true,
		},
		Quantity: int32(arg.Quantity),
	}

	err = s.repo.AddProductToCartTx(ctx, argRepo)
	if err != nil {
		return err
	}

	return nil

}

func (s *service) DeleteProductOnCartById(ctx context.Context, authHeader string, arg validator.DeleteProductOnCart) error {

	claims, err := s.getJWTClaims(authHeader)
	if err != nil {
		return err
	}

	uuidUser, err := uuid.Parse(claims.Issuer)
	if err != nil {
		return err
	}
	uuidCartItem, err := uuid.Parse(arg.CartItemId)
	if err != nil {
		return err
	}

	err = s.repo.DeleteProductOnCartUserTx(ctx, repository.DeleteProductOnCartUserParams{
		UserID:     uuid.NullUUID{UUID: uuidUser, Valid: true},
		CartItemId: uuidCartItem,
	})
	if err != nil {
		return err
	}

	return nil

}
