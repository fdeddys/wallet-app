package controllers

import (
	"net/http"

	"wallet-app/service"
	"wallet-app/validators"

	"github.com/gin-gonic/gin"
)

var walletService = service.WalletService{}

func GetBalance(c *gin.Context) {
	id := c.Param("id")
	user, err := walletService.GetBalance(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User id not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"name": user.Name, "balance": user.Balance})
}

func Withdraw(c *gin.Context) {
	var req validators.WithdrawRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := validators.ValidateWithdraw(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// var user models.User
	user, msg := walletService.Withdraw(req)
	if msg != "" {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": msg})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Withdraw successful", "name": user.Name, "balance": user.Balance})
}
