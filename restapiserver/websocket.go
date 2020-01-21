package main

import (
	"github.com/gorilla/websocket"
	"log"
	"strconv"
)

type Client struct {
	hub *Hub
	websocket *websocket.Conn
	send chan int
}



//Запись в веб-сокет
func (c *Client) websocketWrite() {
	defer func() {
		ticker.Stop()
		c.websocket.Close()
	}()
	for {
		send := <- c.send
		var message []byte = []byte("{\"count\":" + strconv.Itoa(send) + "\"}")
		if err := c.websocket.WriteMessage(1, message); err != nil {
			//TextMessage = 1 || BinaryMessage = 2 || CloseMessage = 8 || PingMessage = 9 || PongMessage = 10
			log.Println(err)
		}
	}
}

//Чтение из веб-сокета
func (c *Client) websocketRead() {
	defer func() {
		c.hub.unregister <- c
		c.websocket.Close()
	}()

}








//if msg == "stop" {
//
//	break
//}
//var count int
//var timeStamp time.Time
//count, timeStamp = service.GetDataFromCounter()

//var message []byte = []byte("{\"count\":" + strconv.Itoa(service.GetDataFromChanel()) + "\"}")

//		var message []byte = []byte("{\"count\":" + strconv.Itoa(<-ch) + ",\"timestamp\":" + "\"" + timeStamp.String() + "\"}")
// Write message back to browser

//if err = websocket.WriteMessage(1, message); err != nil {  //TextMessage = 1 || BinaryMessage = 2 || CloseMessage = 8 || PingMessage = 9 || PongMessage = 10
//	return
