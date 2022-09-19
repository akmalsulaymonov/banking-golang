package card

import "bank/pkg/bank/types"

const withdrawLimit = 20_000_00
const depositLimit = 50_000_00

func Withdraw(card *types.Card, amount types.Money) {
	if card.Active == true && card.Balance >= amount && amount > 0 && amount <= withdrawLimit {
		card.Balance -= amount
	}
}

func Deposit(card *types.Card, amount types.Money) {
	if card.Active == true && amount > 0 && amount <= depositLimit {
		card.Balance += amount
	}
}

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	if card.Active == true && card.Balance > 0 {
		bonus := (int(card.MinBalance) * percent / 100 * daysInMonth) / daysInYear
		if bonus > 5_000_00 {
			bonus = 0
		}
		card.Balance += types.Money(bonus)
	}
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

func Total(cards []types.Card) types.Money {
	sum := types.Money(0)
	for _, card := range cards {
		if card.Active && card.Balance > 0 {
			sum += card.Balance
		}
	}
	return sum
}

func PaymentSources(cards []types.Card) []types.PaymentSource {
	var payment [] types.PaymentSource
	var payment_type types.PaymentSource
	for _, card := range cards {
		if card.Active && card.Balance > 0 {
			payment_type.Type = "Card"
			payment_type.Number = string(card.PAN)
			payment_type.Balance = card.Balance
			payment = append(payment, payment_type)
		}
	}
	return payment
}
