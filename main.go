package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/okmkey45/greedy_golang/greedy"
)

var defaultCoinList = []int{1, 5, 10, 25, 50}

func main() {
	defaultCoinListDescription := fmt.Sprintf("For default coin list (%v), please enter '0'", defaultCoinList)

	rawCoinList := askForUserInput(
		"Enter coin list separated by comma",
		defaultCoinListDescription,
		"Note: coin values must be sorted ascending. If not, unexpected behavior might occur",
	)

	rawChange := askForUserInput("Enter change (only integer values)")

	coinList, err := parseCoinList(rawCoinList)
	if err != nil {
		panic(err)
	}

	change, err := strToInt(rawChange)
	if err != nil {
		panic(err)
	}

	coinsForChange, err := greedy.FindMinimumNumberOfCoins(coinList, change)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Considering coin list: %v\n", coinList)
	fmt.Printf("And change equals to: %d\n", change)

	fmt.Printf("The minimum number of coins for change are: %d\n", coinsForChange)
}

func parseCoinList(rawCoinList string) ([]int, error) {
	if rawCoinList == "0" {
		return defaultCoinList, nil
	}

	coinStrList := strings.Split(rawCoinList, ",")

	coinList := []int{}

	for _, coinStr := range coinStrList {
		coin, err := strToInt(coinStr)
		if err != nil {
			return nil, err
		}

		coinList = append(coinList, coin)
	}

	return coinList, nil
}

func strToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

func askForUserInput(descriptions ...string) string {
	for _, desc := range descriptions {
		fmt.Println(desc)
	}

	var input string
	fmt.Scan(&input)

	fmt.Println()

	return input
}
