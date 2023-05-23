package greedy

import "errors"

// ErrCoinListIsEmpty defines an error when a coin list is empty
// that will be used for finding minimum number of coins for change
var ErrCoinListIsEmpty = errors.New("coin list is empty")
