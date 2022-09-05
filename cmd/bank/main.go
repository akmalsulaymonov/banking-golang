package main

import (
	"bank/pkg/bank/card"
	"bank/pkg/bank/transfer"
	"bank/pkg/bank/types"
	"fmt"
)

func main() {
	amount := types.Money(1000_00)
	bonus := transfer.Bonus(amount)
	fmt.Println(bonus)

	/*
		card := types.Card{
			ID:       123,
			PAN:      "123",
			Balance:  999_99,
			Currency: "TJS",
			Color:    "white",
			Name:     "Infinity",
			Active:   true,
		}

		fmt.Printf("%+v", card)

		id := card.ID
		fmt.Println(id)

		card.Balance = 1000_00
		fmt.Printf("%+v", card)
	*/

	result := card.IssueCard("USD", "Black", "AKMAL SULAYMONOV")
	fmt.Println(result)
}

func handle(card types.Card) {
	// todo
}
