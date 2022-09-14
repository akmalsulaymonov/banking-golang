package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleMax() {
	payments := []types.Payment{
		{
			ID:     103,
			Amount: 3_000_00,
		},
		{
			ID:     104,
			Amount: 17_000_00,
		},
		{
			ID:     105,
			Amount: 9_000_00,
		},
	}
	max := Max(payments)
	fmt.Println(max.ID)
	// Output: 104
}
