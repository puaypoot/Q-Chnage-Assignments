package services

import (
	"context"
	"fmt"
	"q-change-assignment/src/cashier/dao"
)

func (sv *CashierService) ChangeMoney(ctx context.Context, d *dao.ChangeMoneyRequest) (*dao.ChangeMoneyResponse, error) {
	amount := d.CashValue - d.ProductPrice
	if amount < 0 {
		return &dao.ChangeMoneyResponse{}, fmt.Errorf("cash value is less than the product price")
	}

	// GET CHANGE COINS
	changeCoins, err := sv.findMinimumChangeCoins(amount)
	if err != nil {
		return &dao.ChangeMoneyResponse{}, err
	}

	// UPDATE STORED COINS AMOUNT
	storedCoinsMap := make(map[float64]*StoredCoin)
	for _, c := range sv.StoredCoins {
		storedCoinsMap[c.Value] = c
	}

	for _, uc := range changeCoins {
		storedCoinsMap[uc.CoinValue].Amount -= uc.Amount
	}

	// changeCoins := make
	return &dao.ChangeMoneyResponse{
		ChangeCoins: changeCoins,
	}, nil
}
