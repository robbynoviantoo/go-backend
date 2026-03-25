package service

import (
	"errors"

	"go-backend/models"
	"go-backend/repository"
	"go-backend/utils"

	"golang.org/x/crypto/bcrypt"
)

func Register(user models.User) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hash)

	return repository.CreateUser(user)
}

func Login(email, password string) (string, error) {
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("wrong password")
	}

	token, _ := utils.GenerateToken(user.ID)

	return token, nil
}