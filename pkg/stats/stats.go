package stats

import (
	"github.com/bahromzokirov/bank/v2/pkg/types"
)

 
func CategoriesAvg (payments []types.Payment) map[types.Category] types.Money{
	ans := map [types.Category]types.Money{}
	for _, i := range payments {
		ans[i.Category] =TotalInCategory(payments,i.Category)
	}
	return ans
}
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	s := 0
	c:= 0
	for _, i := range payments {
		if i.Category == category{
			s += int(i.Amount)
			c ++ ;
		}
		
	}
	return types.Money(s/c)
}
