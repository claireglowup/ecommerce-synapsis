package service

import (
	"context"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func (s *service) MidtransPayment(ctx context.Context, authHeader string) (*snap.Response, error) {

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

	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var itemMidtrans []midtrans.ItemDetails
	var price, amount int

	for _, each := range cartUser {
		price, err = strconv.Atoi(each.Price)
		if err != nil {
			return nil, err
		}

		priceByQuantity := price * int(each.Quantity)

		seed := midtrans.ItemDetails{
			ID:       each.ID.String(),
			Name:     each.Name,
			Price:    int64(price),
			Category: each.Category,
			Qty:      each.Quantity,
		}

		itemMidtrans = append(itemMidtrans, seed)

		amount += priceByQuantity
	}

	var m = snap.Client{}
	m.New(os.Getenv("SERVER_KEY_MIDTRANS"), midtrans.Sandbox)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  uuid.NewString(),
			GrossAmt: int64(amount),
		},
		Items: &itemMidtrans,
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Name,
			Email: user.Email,
		},
	}

	snapResp, _ := m.CreateTransaction(req)
	if err != nil {
		return nil, err
	}

	return snapResp, nil
}
