package repository

import (
	"context"

	"github.com/google/uuid"
)

type AddProductToCartParams struct {
	UserID    uuid.NullUUID
	ProductID uuid.NullUUID
	Quantity  int32
}

type DeleteProductOnCartUserParams struct {
	UserID     uuid.NullUUID
	CartItemId uuid.UUID
}

func (s *SQLStore) AddProductToCartTx(ctx context.Context, arg AddProductToCartParams) error {

	err := s.ExecTx(ctx, func(q *Queries) error {

		var err error

		idCart, err := q.AddIdUserToCart(ctx, arg.UserID)
		if err != nil {
			return err
		}

		err = q.AddProductToCartItems(ctx, AddProductToCartItemsParams{
			CartID:    uuid.NullUUID{UUID: idCart, Valid: true},
			ProductID: arg.ProductID,
			Quantity:  arg.Quantity,
		})
		if err != nil {
			return err
		}

		return err
	})

	if err != nil {
		return err
	}

	return nil

}

func (s *SQLStore) DeleteProductOnCartUserTx(ctx context.Context, arg DeleteProductOnCartUserParams) error {

	err := s.ExecTx(ctx, func(q *Queries) error {

		var err error

		idCart, err := q.GetCartIdByUserId(ctx, arg.UserID)
		if err != nil {
			return err
		}

		err = q.DeleteProductOnCartById(ctx, DeleteProductOnCartByIdParams{
			ID:     arg.CartItemId,
			CartID: uuid.NullUUID{UUID: idCart, Valid: true},
		})
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		return err
	}

	return nil

}
