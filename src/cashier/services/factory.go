package services

import (
	"context"
	"q-change-assignment/src/cashier/dao"
)

type StoredCoin struct {
	Value   float64
	Maximum int64
	Amount  int64
}

type CashierService struct {
	StoredCoins []*StoredCoin
}

type ICashierService interface {
	ChangeMoney(ctx context.Context, d *dao.ChangeMoneyRequest) (*dao.ChangeMoneyResponse, error)
	AddCoins(ctx context.Context, d *dao.AddCoinRequest) (*dao.AddCoinResponse, []error)
}

func NewCashierService() ICashierService {
	return &CashierService{
		StoredCoins: []*StoredCoin{
			{
				Value:   1000.00,
				Maximum: 10,
				Amount:  10,
			},
			{
				Value:   500.00,
				Maximum: 20,
				Amount:  20,
			},
			{
				Value:   100.00,
				Maximum: 15,
				Amount:  15,
			},
			{
				Value:   50.00,
				Maximum: 20,
				Amount:  20,
			},
			{
				Value:   20.00,
				Maximum: 30,
				Amount:  30,
			},
			{
				Value:   10.00,
				Maximum: 20,
				Amount:  20,
			},
			{
				Value:   5.00,
				Maximum: 20,
				Amount:  20,
			},
			{
				Value:   1.00,
				Maximum: 20,
				Amount:  20,
			},
			{
				Value:   0.25,
				Maximum: 50,
				Amount:  50,
			},
		},
	}
}
