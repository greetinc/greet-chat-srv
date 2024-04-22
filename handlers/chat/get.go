package handlers

import (
	"encoding/json"
	dto "greet-chat-srv/dto/chat"
	"greet-chat-srv/entity"

	util "github.com/greetinc/greet-util/s"

	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type Notification struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *chatHandler) HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	h.addClient <- conn

	for {
		var msg entity.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			h.removeClient <- conn
			return
		}

		h.broadcast <- msg
	}
}

func (h *chatHandler) HandleMessages() {
	for {
		select {
		case client := <-h.addClient:
			h.mutex.Lock()
			h.clients[client] = true
			h.mutex.Unlock()

		case client := <-h.removeClient:
			h.mutex.Lock()
			delete(h.clients, client)
			h.mutex.Unlock()

		case msg := <-h.broadcast:
			if err := h.SendMessage(msg); err != nil {
				log.Println("Error saving message to database:", err)
				continue
			}

			h.mutex.Lock()
			for client := range h.clients {
				err := client.WriteJSON(msg)
				if err != nil {
					log.Println(err)
					client.Close()
					delete(h.clients, client)
				}
			}
			h.mutex.Unlock()
		}
	}
}

func (h *chatHandler) SendMessage(msg entity.Message) error {
	if err := h.serviceChat.SendMessage(&msg); err != nil {
		return err
	}

	return nil
}

func (h *chatHandler) Chat(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println(err)
		return err
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return err
		}

		var receivedMessage entity.Message
		if err := json.Unmarshal(p, &receivedMessage); err != nil {
			log.Println("Error unmarshaling message:", err)
			continue
		}
		receivedMessage.ID = util.GenerateRandomString()

		if err := h.SendMessage(receivedMessage); err != nil {
			log.Println("Error saving message to database:", err)
			continue
		}

		sender_id := receivedMessage.SenderID
		receiver_id := receivedMessage.ReceiverID

		messages, err := h.GetMessagesFromDatabase(sender_id, receiver_id)
		if err != nil {
			log.Println("Error fetching messages from database:", err)
			continue
		}

		for _, msg := range messages {
			byteMsg, err := json.Marshal(msg)
			if err != nil {
				log.Println("Error marshaling message:", err)
				continue
			}

			if err := conn.WriteMessage(messageType, byteMsg); err != nil {
				log.Println(err)
				continue
			}
		}

		notification := Notification{Type: "new_message", Content: "Ada pesan baru!"}
		byteNotification, err := json.Marshal(notification)
		if err != nil {
			log.Println("Error marshaling notification:", err)
			continue
		}

		if err := conn.WriteMessage(messageType, byteNotification); err != nil {
			log.Println(err)
			continue
		}
	}
}

func (h *chatHandler) GetMessagesFromDatabase(sender_id, receiver_id string) ([]dto.ChatResponse, error) {
	messages, err := h.serviceChat.GetAll(sender_id, receiver_id)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
