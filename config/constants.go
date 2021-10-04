package config

const SYMBOL = "btcusdt"
// 20 вместо 15, потому что в документации есть только 3 варианта: 5, 10, 20. Если правильно понял задачу, то количество ордеров в стакане относится к этому (Кол-во ордеров в каждом стакане - 15)
const DEPTH = "20"
const UPDATE_SPEED = "1000ms"
const BASE_URL = "wss://stream.binance.com:9443/ws/"
