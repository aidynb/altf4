package utils

import (
	"log"
	"strconv"
)

// QuantityTotal loops through the data and returns total quantity of the given data
func QuantityTotal(data [][]string) float64 {
	var total float64
	for i := range data {
		quantityStr := data[i][1]

		quantityInt, err := strconv.ParseFloat(quantityStr, 32)
		if err != nil {
			log.Printf("Error converting to int: %s\n", err)
			continue
		}

		total += quantityInt
	}
	return total
}
