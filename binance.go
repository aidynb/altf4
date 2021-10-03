package main

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"

	"github.com/aidynb/altf4/config"
	"github.com/aidynb/altf4/utils"
)

// OrderBook - structure of the order book with symbol, bids and asks
type OrderBook struct {
	Symbol            string     `json:"symbol"`
	Bids              [][]string `json:"bids"`
	Asks              [][]string `json:"asks"`
	TotalQuantityBids float64    `json:"total_quantity_bids"`
	TotalQuantityAsks float64    `json:"total_quantity_asks"`
}

// BinanceWSConnection connects to binance API and sends data to the channel
func BinanceWSConnection(c chan OrderBook) {
	// main endpoint
	endpoint := config.BASE_URL + config.SYMBOL + "@depth" + config.DEPTH + "@" + config.UPDATE_SPEED

	// websocket connection to the binance api
	conn, _, err := websocket.DefaultDialer.Dial(endpoint, nil)
	if err != nil {
		log.Fatalf("Couldn't connect to the Binance API: %s\n", err)
	}

	log.Println("Binance websocket connection established")

	orderBook := OrderBook{Symbol: config.SYMBOL}

	// starting the infinite loop to get the order book from binance api
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error in message read: %s\n", err)
			continue
		}

		// parse the incoming message
		json.Unmarshal([]byte(message), &orderBook)

		// get total quantity of bids and asks
		orderBook.TotalQuantityBids = utils.QuantityTotal(orderBook.Bids)
		orderBook.TotalQuantityAsks = utils.QuantityTotal(orderBook.Asks)

		// send new order book to the channel
		c <- orderBook

	}
}
