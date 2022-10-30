package stats

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/akmalsulaymonov/banking-golang/pkg/bank/types"
)

func TestFilterByCategory_nil(t *testing.T) {
	var payments []types.Payment
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len not equal 0")
	}
}

func TestFilterByCategory_empty(t *testing.T) {
	payments := []types.Payment{}
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len not equal 0")
	}
}

func TestFilterByCategory_notFound(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "fun"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "game"},
	}
	result := FilterByCategory(payments, "mobile")
	if len(result) != 0 {
		t.Error("result len not equal 0")
	}
}

func TestFilterByCategory_foundOne(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "fun"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "game"},
	}

	expected := []types.Payment{
		{ID: 2, Category: "food"},
	}

	result := FilterByCategory(payments, "food")

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestFilterByCategory_foundMultiple(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "fun"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "game"},
		{ID: 6, Category: "auto"},
	}

	expected := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 6, Category: "auto"},
		{ID: 4, Category: "auto"},
	}

	result := FilterByCategory(payments, "auto")

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestCategoriesTotal(t *testing.T) {
	payments := []types.Payment{
		{
			ID:       103,
			Amount:   3_000_00,
			Category: "Fun",
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

	expected := map[types.Category]types.Money{
		"Auto": 20_000_00,
		"Shop": 30_000_00,
		"Fun":  3_000_00,
	}

	result := CategoriesTotal(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   10,
			Category: "auto",
		},
		{
			ID:       2,
			Amount:   50,
			Category: "auto",
		},
		{
			ID:       3,
			Amount:   10,
			Category: "food",
		},
		{
			ID:       4,
			Amount:   30,
			Category: "food",
		},
		{
			ID:       4,
			Amount:   10,
			Category: "fun",
		},
	}

	expected := map[types.Category]types.Money{
		"auto": 30,
		"food": 20,
		"fun":  10,
	}

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Invalid match result map, expected: %v,actual: %v", expected, result)
	}
}

func TestPeriodDynamic(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}
	second := map[types.Category]types.Money{
		"auto": 5,
		"food": 3,
		"mobile":  10,
	}
	expected := map[types.Category]types.Money{
		"auto": -5,
		"food": -17,
		"mobile":  10,
	}
	result := PeriodsDynamic(first, second)

	for key := range result {
		fmt.Println(key, ":", result[key])
	}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Invalid match result map, expected: %v,actual: %v", expected, result)
	}
}
