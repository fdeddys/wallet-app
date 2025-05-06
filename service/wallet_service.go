package service

import (
	"wallet-app/config"
	"wallet-app/models"
	"wallet-app/validators"
)

type WalletService struct{}

func (s *WalletService) GetBalance(id string) (models.User, error) {
	var user models.User
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *WalletService) Withdraw(req validators.WithdrawRequest) (models.User, string) {
	var user models.User

	if err := config.DB.Where("id = ?", req.ID).First(&user).Error; err != nil {
		return models.User{}, "User not found"
	}

	if user.Balance < req.Amount {
		return models.User{}, "Insufficient balance"
	}

	user.Balance -= req.Amount
	config.DB.Save(&user)

	return user, ""
}
