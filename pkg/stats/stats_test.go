package stats

import (
	"reflect"
	"testing"

	"github.com/bahromzokirov/bank/pkg/types"
)

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{ Category: "auto", Amount: 10_000},
		{ Category: "auto", Amount: 20_000},
		{ Category: "auto", Amount: 30_000},
	}
	expected := map[types.Category]types.Money{
		"auto": 20_000,
	}
	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected. %v, actual: %v", expected, result)
	}
}
