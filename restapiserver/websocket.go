package main

import (
	"fmt"
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
		c.websocket.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.websocket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			var msg []byte = []byte("{\"count\":\"" + strconv.Itoa(message) + "\"}")

			fmt.Printf("send: {\"count\":\"" + strconv.Itoa(message) + "\"}\n")	//Сообщение

			if err := c.websocket.WriteMessage(1, msg); err != nil {
				//TextMessage = 1 || BinaryMessage = 2 || CloseMessage = 8 || PingMessage = 9 || PongMessage = 10
				log.Println(c.websocket.LocalAddr().String() + " - Exit")
				c.hub.unregister <- c
			}
		}
	}
}

//Чтение из веб-сокета
func (c *Client) websocketRead() {
	//defer func() {
	//	c.hub.unregister <- c
	//	c.websocket.Close()
	//}()
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
