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
