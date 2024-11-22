package service

import (
	"context"
	"errors"
	"os"
	"synapsis-ecommerce/src/helper/validator"
	"synapsis-ecommerce/src/repository"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func (u *service) Register(ctx context.Context, user validator.UserRegister) error {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
	if err != nil {
		return err
	}

	newUser := repository.RegisterUserParams{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(passwordHash),
	}

	err = u.repo.RegisterUser(ctx, newUser)
	if err != nil {
		return err
	}

	return nil

}

func (u *service) Login(ctx context.Context, user validator.UserLogin) (string, error) {

	data, err := u.repo.Login(ctx, user.Email)
	if err != nil {
		return "", errors.New("Email is not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(user.Password))
	if err != nil {
		return "", errors.New("Wrong password")

	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    data.ID.String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})

	token, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return token, nil

}
