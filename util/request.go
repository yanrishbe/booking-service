package util

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/yanrishbe/booking-service/model"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	ID         string
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
}

type AccountRequest struct {
	ID     string `json:"id,omitempty"`
	Bank   string `json:"bank"`
	Amount string `json:"amount"`
}

func AccountFromRequest(req AccountRequest) (*model.Account, error) {
	acc := model.Account{
		ID:   req.ID,
		Bank: req.Bank,
	}
	amount := strings.Split(req.Amount, ",")
	var cents int
	var err error
	if len(amount) == 0 {
		return &acc, nil
	}
	if len(amount) == 2 {
		cents, err = strconv.Atoi(amount[1])
		if err != nil {
			return nil, fmt.Errorf("could not convert cents: %v", err)
		}
	}
	main, err := strconv.Atoi(amount[0])
	if err != nil {
		return nil, fmt.Errorf("could not convert main: %v", err)
	}
	acc.Amount = main*100 + cents
	return &acc, nil
}
