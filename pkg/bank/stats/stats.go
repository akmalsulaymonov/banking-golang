package stats

import "github.com/akmalsulaymonov/banking-golang/pkg/bank/types"

func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)
	for _, payment := range payments {
		sum += payment.Amount
	}
	return sum / types.Money(len(payments))
}
