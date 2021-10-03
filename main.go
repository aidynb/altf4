package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

const BASE_URL = "wss://stream.binance.com:9443/ws/btcusdt@depth"

func main() {
	conn, _, err := websocket.DefaultDialer.Dial(BASE_URL, nil)

	// fmt.Println("conn: ", conn)
	// fmt.Println("res: ", res)
	// fmt.Println("err: ", err)

	if err != nil {
		log.Fatalln(err)
	}

	for {
		msgType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error in message read: ", err)
			return
		}

		fmt.Printf("message type: %d, message: %s\n", msgType, message)
	}
}
