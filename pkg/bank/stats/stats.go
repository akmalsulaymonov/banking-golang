package stats

import "github.com/akmalsulaymonov/banking-golang/pkg/bank/types"

func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)
	for _, payment := range payments {
		sum += payment.Amount
	}
	return sum / types.Money(len(payments))
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)
	for _, payment := range payments {
		if category == payment.Category {
			sum += payment.Amount
		}
	}
	return sum
}