package greedy

// FindMinimumNumberOfCoins defines the minimum number of coins needed
// for a specific change
//
// NOTES:
//
// 1. This function assumes that the coin list is sorted ascending
// 2. For some scenarios might return a correct answer but NOT the best one
//
// TODO: add/improve implementation to handle special cases and get the best output
func FindMinimumNumberOfCoins(
	coinList []int,
	change int,
) (int, error) {
	if change <= 0 {
		// if change equals to zero,
		// no need to continue
		// just return current coins for change
		return 0, nil
	}

	if len(coinList) <= 0 {
		// if list of coins is empty,
		// there is no way to perform proper comparison
		// just return current coins for change
		return 0, ErrCoinListIsEmpty
	}

	coinsForChange := 0

	// a variable to hold remaining change is needed
	// do not modify original value
	remainingChange := change

	for i := len(coinList) - 1; i >= 0; i-- {
		// assuming last value of the list is the biggest coin
		biggestCoin := coinList[i]

		for remainingChange >= biggestCoin {
			coinsForChange++
			remainingChange -= biggestCoin
		}
	}

	return coinsForChange, nil
}
