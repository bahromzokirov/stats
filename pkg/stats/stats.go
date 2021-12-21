package stats

import (
	"github.com/bahromzokirov/bank/pkg/types"
)

func Avg(payments []types.Payment) types.Money {
	amount := (0)
	id := (0)
	for _, i := range payments {
		amount += int(i.Amount)
		id++
	}

	return types.Money(amount / id)
}
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	c:= 0
	for _, i := range payments {
		if i.Category == category{
			c += int(i.Amount)
		}
		
	}
	return types.Money(c)
}
