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
			Category: "Chemist",
		},
		{
			ID:       105,
			Amount:   10_000_00,
			Category: "Restaurant",
		},
	}
	avgSum := Avg(payments)
	fmt.Println(avgSum)
	// Output: 1000000
}
