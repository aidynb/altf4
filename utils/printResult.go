package utils

import (
	"fmt"
)

// PrintResult prints result to the console
func PrintResult(symbol string, bids, asks [][]string, totalBids, totalAsks float64) {
	fmt.Printf("Order Book: %s\n", symbol)

	fmt.Println("BID:")

	for i := range bids {
		price := bids[i][0]
		quantity := bids[i][1]
		fmt.Printf("Price: %s\tQuantity: %s\n", price, quantity)
	}
	fmt.Println("ASK:")

	for i := range asks {
		price := asks[i][0]
		quantity := asks[i][1]
		fmt.Printf("Price: %s\tQuantity: %s\n", price, quantity)
	}

	fmt.Printf("Total BID quantity: %f\n", totalBids)
	fmt.Printf("Total ASK quantity: %f\n", totalAsks)
}
