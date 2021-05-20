package services

import (
	"context"
	"fmt"
	"q-change-assignment/src/cashier/dao"
)

func (sv *CashierService) AddCoins(ctx context.Context, d *dao.AddCoinRequest) (*dao.AddCoinResponse, []error) {
	// GET CURRENT STORED COINS
	storedCoinsMap := make(map[float64]*StoredCoin)
	for _, c := range sv.StoredCoins {
		storedCoinsMap[c.Value] = c
	}

	// VALIDATE COINS AMOUNT
	errs := make([]error, 0, len(d.AddedCoins))
	for _, ac := range d.AddedCoins {
		v, found := storedCoinsMap[ac.CoinValue]
		if !found {
			errs = append(errs, fmt.Errorf("coin value: %v not support", ac.CoinValue))
			continue
		}

		if ac.Amount+v.Amount > v.Maximum {
			errs = append(errs, fmt.Errorf("coin value: %v total (%v) coins is more than maximum (%v)", ac.CoinValue, ac.Amount+v.Amount, v.Maximum))
			continue
		}
	}
	if len(errs) > 0 {
		return &dao.AddCoinResponse{}, errs
	}

	// UPDATE STORED COINS AMOUNT
	for _, ac := range d.AddedCoins {
		storedCoinsMap[ac.CoinValue].Amount += ac.Amount
	}

	// GEN UPDATE RESULT
	result := make([]dao.Coin, 0, len(sv.StoredCoins))
	for _, c := range sv.StoredCoins {
		result = append(result, dao.Coin{
			CoinValue: c.Value,
			Amount:    c.Amount,
		})
	}

	return &dao.AddCoinResponse{
		StoredCoins: result,
	}, nil
}
