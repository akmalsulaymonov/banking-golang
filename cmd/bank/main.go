package main

import (
	"bank/pkg/bank/types"
	"bank/pkg/bank/transfer"
	"fmt"
)

func main() {
	amount := types.Money(1000_00)
	bonus := transfer.Bonus(amount)
	fmt.Println(bonus)
}
