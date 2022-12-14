package main

import (
	"fmt"

	"github.com/akmalsulaymonov/banking-golang/pkg/bank/card"
	"github.com/akmalsulaymonov/banking-golang/pkg/bank/transfer"
	"github.com/akmalsulaymonov/banking-golang/pkg/bank/types"
)

func main() {
	amount := types.Money(1000_00)
	bonus := transfer.Bonus(amount)
	fmt.Println(bonus) // 500

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

	//result := card.IssueCard("USD", "Black", "AKMAL SULAYMONOV")
	//fmt.Println(result)

	//cardNew := types.Card{Balance: 30_000_00, Active: true}
	//res := card.Withdraw(&cardNew, 10_000_00)
	//fmt.Println(res) // 20_000_00

	balance := int64(60_000_00)
	withdraw(&balance, 10_000_00)
	fmt.Println(balance) // 50_000_00

	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active:  true,
			PAN:     "5058 xxxx xxxx 0000",
		},
		{
			Balance: 50_000_00,
			Active:  true,
			PAN:     "5058 xxxx xxxx 1111",
		},
		{
			Balance: 10_000_00,
			Active:  false,
			PAN:     "5058 xxxx xxxx 2222",
		},
		{
			Balance: -20_000_00,
			Active:  true,
			PAN:     "5058 xxxx xxxx 3333",
		},
	}
	fmt.Println(card.PaymentSources(cards))

}

func withdraw(balance *int64, amount int64) {
	*balance -= amount
}

func handle(card types.Card) {
	// todo
}
