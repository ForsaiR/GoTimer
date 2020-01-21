package main

type Hub struct {
	clients map[*Client]bool

	broadcast chan int

	register chan *Client

	unregister chan *Client
}

//Создание нового хаба
func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan int),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

//Запуск работы хаба
func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		//case message := <-h.broadcast:
		//	for client := range h.clients {
		//		select {
		//		case client.send <- message:
		//		default:
		//			close(client.send)
		//			delete(h.clients, client)
		//		}
		//	}
		}
	}
}
