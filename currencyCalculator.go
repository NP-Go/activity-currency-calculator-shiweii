package main

import (
	"fmt"
	"math"
	"strconv"
)

func CurrencyCalculator(oneDollar, fiftyCent, twentyCent, tenCent, fiveCent float64) (float64, float64, float64) {
	//Insert your code here
	totalAmount := oneDollar + (0.5 * fiftyCent) + (0.2 * twentyCent) + (0.1 * tenCent) + (0.05 * fiveCent)
	changeAmount := math.Mod(totalAmount, 2)
	twoDollarNotes := (totalAmount - changeAmount) / 2

	//Do not remove this
	fmt.Println("Total:", totalAmount, "Two Dollar Notes:", twoDollarNotes, " Change: ", changeAmount)
	return totalAmount, twoDollarNotes, changeAmount
}

// Read user input and covert string to integer, question will be repeated if non integer is entered.
func inputCheck(m string, i string) float64 {
	var ret int
	for {
		fmt.Println(m)
		fmt.Scanln(&i)
		value, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println("Please enter a valid number of coins.")
		} else {
			ret = value
			break
		}
	}
	return float64(ret)
}

func main() {
	var oneDollarCoin float64
	var fiftyCentCoin float64
	var twentyCentCoin float64
	var tenCentCoin float64
	var fiveCentCoin float64
	var oneDollarCoinInput, fiftyCentCoinInput, twentyCentCoinInput, tenCentCoinInput, fiveCentCoinInput string

	oneDollarCoin = inputCheck("Enter the number of 1-dollar coins: ", oneDollarCoinInput)
	fiftyCentCoin = inputCheck("Enter the number of 50-cent coins: ", fiftyCentCoinInput)
	twentyCentCoin = inputCheck("Enter the number of 20-cent coins: ", twentyCentCoinInput)
	tenCentCoin = inputCheck("Enter the number of 10-cent coins: ", tenCentCoinInput)
	fiveCentCoin = inputCheck("Enter the number of 5-cent coins: ", fiveCentCoinInput)

	CurrencyCalculator(oneDollarCoin, fiftyCentCoin, twentyCentCoin, tenCentCoin, fiveCentCoin)
}
