package dao

type ChangeMoneyRequest struct {
	CashValue    float64 `json:"cash_value"`
	ProductPrice float64 `json:"product_price"`
}

type ChangeMoneyResponse struct {
	ChangeCoins []Coin `json:"change_coins"`
}
