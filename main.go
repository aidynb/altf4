package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/gorilla/websocket"
)

const SYMBOL = "btcusdt"
const DEPTH = "20"
const UPDATE_SPEED = "1000ms"
const BASE_URL = "wss://stream.binance.com:9443/ws/"

// OrderBook - structure of the order book with symbol, bids and asks
type OrderBook struct {
	Symbol            string
	Bids              [][]string `json:"bids"`
	Asks              [][]string `json:"asks"`
	TotalQuantityBids float64
	TotalQuantityAsks float64
}

// PrintResult prints result to the console
func PrintResult(orderBook *OrderBook) {
	fmt.Printf("Order Book: %s\n", orderBook.Symbol)

	fmt.Println("BID:")

	for i := range orderBook.Bids {
		price := orderBook.Bids[i][0]
		quantity := orderBook.Bids[i][1]
		fmt.Printf("Price: %s\tQuantity: %s\n", price, quantity)
	}
	fmt.Println("ASK:")

	for i := range orderBook.Asks {
		price := orderBook.Asks[i][0]
		quantity := orderBook.Asks[i][1]
		fmt.Printf("Price: %s\tQuantity: %s\n", price, quantity)
	}

	fmt.Printf("Total BID quantity: %f\n", orderBook.TotalQuantityBids)
	fmt.Printf("Total ASK quantity: %f\n", orderBook.TotalQuantityAsks)
}

// QuantityTotal loops through the data and returns total quantity of the given data
func QuantityTotal(data [][]string) float64 {
	var total float64
	for i := range data {
		quantityStr := data[i][1]

		quantityInt, err := strconv.ParseFloat(quantityStr, 32)
		if err != nil {
			fmt.Printf("Error converting to int: %s\n", err)
		}

		total += quantityInt
	}
	return total
}

func main() {
	// main endpoint
	endpoint := BASE_URL + SYMBOL + "@depth" + DEPTH + "@" + UPDATE_SPEED

	// websocket connection to the binance api
	conn, _, err := websocket.DefaultDialer.Dial(endpoint, nil)
	if err != nil {
		log.Fatalln(err)
	}

	orderBook := OrderBook{Symbol: SYMBOL}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error in message read: ", err)
			return
		}

		// parse the incoming message
		json.Unmarshal([]byte(message), &orderBook)

		// get total quantity of bids and asks
		orderBook.TotalQuantityBids = QuantityTotal(orderBook.Bids)
		orderBook.TotalQuantityAsks = QuantityTotal(orderBook.Asks)

		PrintResult(&orderBook)

	}
}
