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

func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}

	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
	}

	return categories
}

func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category == category {
			filtered = append(filtered, payment)
		}
	}
	return filtered
}

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	Avg := make(map[types.Category]types.Money)
	index := make(map[types.Category]types.Money)

	for _, i := range payments {
		_, key := Avg[i.Category]
		if key {
			Avg[i.Category] += i.Amount
			index[i.Category]++
		} else {
			Avg[i.Category] = i.Amount
			index[i.Category] = 1
		}
	}
	for key, value := range Avg {
		Avg[key] = value / index[key]
	}
	return Avg
}

func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {
	result := map[types.Category]types.Money{}
	if len(first) > len(second) {
		for key := range first {
			result[key] = second[key] - first[key]
		}
		return result
	}

	for key := range second {
		result[key] = second[key] - first[key]
	}
	return result

}
