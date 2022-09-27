package payment

import "github.com/akmalsulaymonov/banking-golang/pkg/bank/types"

func Max(payments []types.Payment) types.Payment {
	max := payments[0]
	for _, payment := range payments {
		if payment.Amount >= max.Amount {
			max = payment
		}
	}
	return max
}
