package card

import "bank/pkg/bank/types"

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

func Withdraw(card types.Card, amount types.Money) types.Card {
	// TODO: произвести операøии с картой
	if card.Active == true && card.Balance >= amount && amount > 0 && amount <= 20_000_00 {
		card.Balance -= amount
	}
	return card
}
