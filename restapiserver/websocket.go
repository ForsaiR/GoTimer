package main

import (
	"bytes"
	"encoding/json"
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

type WebsocketCommand struct {
	Command string
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

			jsonMessage, err := json.Marshal(*getInstance())
			if err != nil {
				fmt.Println(err)
			}

			var msg []byte = bytes.TrimSpace(bytes.Replace(jsonMessage, []byte{'\n'}, []byte{' '}, -1))

			if err := c.websocket.WriteMessage(1, msg); err != nil {
				//TextMessage = 1 || BinaryMessage = 2 || CloseMessage = 8 || PingMessage = 9 || PongMessage = 10
				log.Println(c.websocket.LocalAddr().String() + " - Exit by send")
				c.hub.unregister <- c
				return
			}

			fmt.Printf("send: {\"count\":\"" + strconv.Itoa(message) + "\"}\n")	//Сообщение
		}
	}
}

//Чтение из веб-сокета
func (c *Client) websocketRead() {
	//{"command":"reset"}
	for {
		_, message, err := c.websocket.ReadMessage()
		if err != nil {
			return
		}

		var jsonMessage WebsocketCommand
		json.Unmarshal(message, &jsonMessage)

		if jsonMessage.Command == "reset" {
			c.hub.resetCounter()
		}
	}
	//defer func() {
	//	c.hub.unregister <- c
	//	c.websocket.Close()
	//}()
}
