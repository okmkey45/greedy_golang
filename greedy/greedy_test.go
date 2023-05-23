package greedy_test

import (
	"testing"

	"github.com/okmkey45/greedy_golang/greedy"
)

func TestBasicFindMinimumNumberOfCoins(t *testing.T) {
	specs := []struct {
		name           string
		coins          []int
		change         int
		expectedAmount int
	}{
		{
			name:   "initial sample",
			coins:  []int{1, 5, 10, 25, 50},
			change: 63,
			// Initial sample describes next output with six coins: [25, 25, 10, 1, 1, 1]
			// Current implementation outputs a better one: [50, 10, 1, 1, 1]
			// one coin less
			expectedAmount: 5,
		},
		{
			name:   "imperfect scenario",
			coins:  []int{1, 5, 6, 9},
			change: 11,
			// NOTE: a better output should be 2 coins: 6 + 5 = 11
			expectedAmount: 3,
		},
	}

	for _, spec := range specs {
		coinsForChange, err := greedy.FindMinimumNumberOfCoins(spec.coins, spec.change)

		if err != nil {
			t.Fatalf("error expected to be nil")
		}

		if coinsForChange != spec.expectedAmount {
			t.Fatalf("coins for change %d expected to be %d", coinsForChange, spec.expectedAmount)
		}
	}
}
