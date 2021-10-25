package main

type transaction struct {
	FromWallet string  `json:"fromWallet"`
	ToWallet   string  `json:"toWallet"`
	Amount     float64 `json:"amount"`
}

func newTransaction(fromWallet, toWallet string, amount float64) (trans transaction) {
	trans = transaction{FromWallet: fromWallet, ToWallet: toWallet, Amount: amount}
	return
}
