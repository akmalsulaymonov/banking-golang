package transfer

import "github.com/akmalsulaymonov/banking-golang/pkg/bank/types"

const bonusPercent = 0.005

func Bonus(amount types.Money) types.Money {
	bonus := types.Money(float64(amount) * bonusPercent)
	return bonus
}

func Total(amount types.Money) types.Money {
	total := amount + Bonus(amount)
	return total
}
