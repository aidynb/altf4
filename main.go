package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/gorilla/websocket"
)

const BASE_URL = "wss://stream.binance.com:9443/ws/btcusdt@depth20@1000ms"

type OrderBook struct {
	Symbol string
	Bids   [][]string `json:"bids"`
	Asks   [][]string `json:"asks"`
}

func main() {
	conn, _, err := websocket.DefaultDialer.Dial(BASE_URL, nil)

	// fmt.Println("conn: ", conn)
	// fmt.Println("res: ", resp)
	// fmt.Println("err: ", err)

	if err != nil {
		log.Fatalln(err)
	}

	var orderBook OrderBook

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error in message read: ", err)
			return
		}

		json.Unmarshal([]byte(message), &orderBook)

		var totalQuantityBid float64
		var totalQuantityAsk float64

		for i := range orderBook.Bids {
			quantityStr := orderBook.Bids[i][1]

			quantityInt, err := strconv.ParseFloat(quantityStr, 32)
			if err != nil {
				fmt.Printf("Error converting to int: %s\n", err)
			}

			totalQuantityBid += quantityInt
		}

		for i := range orderBook.Asks {
			quantityStr := orderBook.Asks[i][1]

			quantityInt, err := strconv.ParseFloat(quantityStr, 32)
			if err != nil {
				fmt.Printf("Error converting to int: %s\n", err)
			}

			totalQuantityAsk += quantityInt
		}

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

		fmt.Printf("Total BID quantity: %f\n", totalQuantityBid)
		fmt.Printf("Total ASK quantity: %f\n", totalQuantityAsk)

	}
}
