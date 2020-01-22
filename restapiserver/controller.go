package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"time"
)

//Стартовая страница
func startPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

//POST запрос для запуска
func putStartCounter(hub *Hub, w http.ResponseWriter, r *http.Request) {
	hub.startCounter()
	json.NewEncoder(w).Encode(map[string]bool{"status": counterStatus()})
}

//POST запрос для остановки счетчика
func putStopCounter(hub *Hub, w http.ResponseWriter, r *http.Request) {
	hub.stopCounter()
	json.NewEncoder(w).Encode(map[string]bool{"status": counterStatus()})
}

//POST запрос для установки значения счетчика 0 (перезапеск счетчика)
func putResetCounter(hub *Hub, w http.ResponseWriter, r *http.Request) {
	hub.resetCounter()
	json.NewEncoder(w).Encode(map[string]bool{"status": counterStatus()})
}

//GET запрос для отправки значения счетчика
func getCounterValue(w http.ResponseWriter, r *http.Request) {
	var count int
	var timeStamp time.Time
	count, timeStamp = (*getInstance()).getDataFromCounter()
	json.NewEncoder(w).Encode(map[string]string{"count": strconv.Itoa(count), "timestamp": timeStamp.String()})
}

//Соединение с веб-сокетом
func serveWebsocket(hub *Hub, w http.ResponseWriter, r *http.Request)  {
	websocket, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{hub: hub, websocket: websocket, send: make(chan int)}
	client.hub.register <- client

	go client.websocketWrite()
	go client.websocketRead()
}