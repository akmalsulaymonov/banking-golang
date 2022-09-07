package card

import "bank/pkg/bank/types"

const withdrawLimit = 20_000_00

func Withdraw(card *types.Card, amount types.Money) types.Card {
	// TODO: произвести операции с картой
	if card.Active == true && card.Balance >= amount && amount > 0 && amount <= withdrawLimit {
		card.Balance -= amount
	}
	return *card
}

func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Name:     name,
		Color:    color,
		Active:   true,
	}
	return card
}
