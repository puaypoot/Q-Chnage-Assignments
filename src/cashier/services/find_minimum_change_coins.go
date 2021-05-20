package services

import (
	"fmt"
	"math"
	"q-change-assignment/src/cashier/dao"
)

// GREEDY
func (sv *CashierService) findMinimumChangeCoins(amount float64) ([]dao.Coin, error) {
	changeCoins := make([]dao.Coin, 0)
	current := amount
	for _, coinInfo := range sv.StoredCoins {
		if current >= coinInfo.Value && coinInfo.Amount > 0 {
			take := math.Min(math.Floor(current/coinInfo.Value), float64(coinInfo.Amount))
			changeCoins = append(changeCoins, dao.Coin{
				CoinValue: coinInfo.Value,
				Amount:    int64(take),
			})
			current = current - take*coinInfo.Value
		}
	}

	if current > 0.00 {
		return changeCoins, fmt.Errorf("total stored coin not enough")
	}

	return changeCoins, nil
}
