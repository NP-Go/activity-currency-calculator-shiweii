package main

import (
	"fmt"
)

func CurrencyCalculator(oneDollar, fiftyCent, twentyCent, tenCent, fiveCent float64) (float64, float64, float64) {
	//Insert your code here

	//Do not remove this
	fmt.Println("Total:", totalAmount, "Two Dollar Notes:", twoDollarNotes, " Change: ", changeAmount)
	return totalAmount, twoDollarNotes, changeAmount
}

func main() {
	var oneDollarCoin float64
	var fiftyCentCoin float64
	var twentyCentCoin float64
	var tenCentCoin float64
	var fiveCentCoin float64

	fmt.Println("Enter the number of 1-dollar coins: ")
	fmt.Scanln(&oneDollarCoin)
	fmt.Println("Enter the number of 50-cent coins: ")
	fmt.Scanln(&fiftyCentCoin)
	fmt.Println("Enter the number of 20-cent coins: ")
	fmt.Scanln(&twentyCentCoin)
	fmt.Println("Enter the number of 10-cent coins: ")
	fmt.Scanln(&tenCentCoin)
	fmt.Println("Enter the number of 5-cent coins: ")
	fmt.Scanln(&fiveCentCoin)
	CurrencyCalculator(oneDollarCoin, fiftyCentCoin, twentyCentCoin, tenCentCoin, fiveCentCoin)
}
