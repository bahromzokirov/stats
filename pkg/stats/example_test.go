package stats

import (
	"github.com/bahromzokirov/bank/pkg/types"
	"fmt"
)

func ExampleTotalInCategory() {
	payment := []types.Payment{
		{
			ID: 1,
			Amount: 10_000,
			Category: "car",
		},
		{
			ID: 3,
			Amount: 20_000,
			Category: "car",
		},
	}
	result := TotalInCategory(payment, "car")
	fmt.Println(result)
// Output: 30000	
}
