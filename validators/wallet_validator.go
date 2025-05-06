package validators

import "errors"

type WithdrawRequest struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
}

func ValidateWithdraw(req WithdrawRequest) error {
	if req.ID == "" {
		return errors.New("id must not be empty")
	}

	if req.Amount <= 0 {
		return errors.New("withdrawal amount must be greater than zero")
	}
	return nil
}
