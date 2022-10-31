package stats

import (
	"fmt"

	"github.com/akmalsulaymonov/banking-golang/pkg/bank/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       103,
			Amount:   3_000_00,
			Category: "Shop",
		},
		{
			ID:       104,
			Amount:   17_000_00,
			Category: "Auto",
		},
		{
			ID:       105,
			Amount:   10_000_00,
			Category: "Shop",
		},
		{
			ID:       106,
			Amount:   20_000_00,
			Category: "Shop",
		},
		{
			ID:       107,
			Amount:   3_000_00,
			Category: "Auto",
		},
	}
	avgSum := Avg(payments)
	total := TotalInCategory(payments, "Auto")
	fmt.Println(avgSum, total)
	// Output: 1060000 2000000
}
