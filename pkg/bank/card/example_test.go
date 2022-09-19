package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleWithdraw_positive() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Withdraw(&card, 10_000_00)
	fmt.Println(card.Balance)
	// Output: 1000000
}

func ExampleWithdraw_noMoney() {
	card := types.Card{Balance: 0, Active: true}
	Withdraw(&card, 10_000_00)
	fmt.Println(card.Balance)
	// Output: 0
}

func ExampleWithdraw_inactive() {
	card := types.Card{Balance: 20_000_00, Active: false}
	Withdraw(&card, 10_000_00)
	fmt.Println(card.Balance)
	// Output: 2000000
}

func ExampleWithdraw_limit() {
	card := types.Card{Balance: 40_000_00, Active: true}
	Withdraw(&card, 30_000_00)
	fmt.Println(card.Balance)
	// Output: 4000000
}

func ExampleCard() {
	card := types.Card{Balance: 20_000_00, Active: true}
	clone := card
	card.Balance -= 10_000_00
	fmt.Println(clone.Balance)
	// Output: 2000000
}

func ExampleDeposit_active() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Deposit(&card, 10_000_00)
	fmt.Println(card.Balance)
	// Output: 3000000
}

func ExampleDeposit_inactive() {
	card := types.Card{Balance: 20_000_00, Active: false}
	Deposit(&card, 10_000_00)
	fmt.Println(card.Balance)
	// Output: 2000000
}

func ExampleDeposit_negative() {
	card := types.Card{Balance: -10_000_00, Active: true}
	Deposit(&card, 30_000_00)
	fmt.Println(card.Balance)
	// Output: 2000000
}

func ExampleDeposit_limit() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Deposit(&card, 60_000_00)
	fmt.Println(card.Balance)
	// Output: 2000000
}

func ExampleAddBonus_positive() {
	result := types.Card{Balance: 10_000_00, MinBalance: 10_000_00, Active: true}
	AddBonus(&result, 3, 30, 365)
	fmt.Println(result.Balance)
	// Output:
	// 1002465
}
func ExampleAddBonus_inactive() {
	result := types.Card{Balance: 10_000_00, MinBalance: 10_000_00, Active: false}
	AddBonus(&result, 3, 30, 365)
	fmt.Println(result.Balance)
	// Output:
	// 1000000
}
func ExampleAddBonus_noMoney() {
	result := types.Card{Balance: -10, MinBalance: 10_000_00, Active: true}
	AddBonus(&result, 3, 30, 365)
	fmt.Println(result.Balance)
	// Output:
	// -10
}

func ExampleAddBonus_limit() {
	result := types.Card{Balance: 10_000_00, MinBalance: 1_000_000_000_00, Active: true}
	AddBonus(&result, 3, 30, 365)
	fmt.Println(result.Balance)
	// Output:
	// 1000000
}

func ExampleTotal() {
	/*
		cards := []types.Card{
			{
				Balance: 10_000_00,
				Active:  true,
			},
			{
				Balance: 10_000_00,
				Active:  false,
			},
			{
				Balance: -20_000_00,
				Active:  true,
			},
			{
				Balance: 10_000_00,
				Active:  true,
			},
		}
		fmt.Println(Total(cards))
		// Output: 2000000
	*/

	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
		},
	}))

	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
		},
		{
			Balance: 2_000_00,
			Active:  true,
		},
	}))

	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  false,
		},
	}))

	fmt.Println(Total([]types.Card{
		{
			Balance: -1_000_00,
			Active:  true,
		},
	}))

	// Output:
	// 100000
	// 300000
	// 0
	// 0
}

func ExamplePaymentSources() {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active:  true,
			PAN:     "5058 xxxx xxxx 0000",
		},
		{
			Balance: 50_000_00,
			Active:  true,
			PAN:     "5058 xxxx xxxx 7777",
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
	payment := PaymentSources(cards)
	for _, card := range payment {
		fmt.Println(card.Number)
	}

	// Output:
	// 5058 xxxx xxxx 0000
	// 5058 xxxx xxxx 7777
}
