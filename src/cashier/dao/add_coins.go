package dao

type AddCoinRequest struct {
	AddedCoins []Coin `json:"added_coins"`
}

type AddCoinResponse struct {
	StoredCoins []Coin `json:"stored_coins"`
}
