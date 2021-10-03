package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// StartServer starts the server
func StartServer(c chan OrderBook) {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("Error upgrading connection: %s\n", err)
			return
		}

		defer conn.Close()

		log.Println("Client connected...")

		go writer(conn, c)

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Printf("Message read error: %s\n", err)
				break
			}

			log.Printf("received: %s\n", message)
		}

	})

	// serve the index template
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/index.html")
	})

	log.Println("Starting server on port 8080...")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

// writer reads from channel and sends to the websocket connection
func writer(conn *websocket.Conn, c chan OrderBook) {
	for {
		orderBook, ok := <-c
		if !ok {
			log.Println("not ok")
		}

		data, _ := json.Marshal(orderBook)

		if err := conn.WriteMessage(1, data); err != nil {
			log.Println("Error in message write: ", err)
		}
	}
}
