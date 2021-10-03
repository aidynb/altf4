package main

func main() {

	// making channel for incoming order book
	c := make(chan OrderBook)

	// starting the server in goroutine
	go StartServer(c)

	// connecting to binance api
	BinanceWSConnection(c)

}
