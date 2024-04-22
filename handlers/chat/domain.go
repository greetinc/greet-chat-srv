package handlers

import (
	"greet-chat-srv/entity"
	s "greet-chat-srv/services/chat"
	"net/http"
	"sync"

	m "github.com/greetinc/greet-middlewares/middlewares"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	// SendMessage(c echo.Context) error
	HandleConnections(w http.ResponseWriter, r *http.Request)
	GetMessagesBetweenUsers(c echo.Context) error
	HandleMessages()
}

type chatHandler struct {
	serviceChat  s.ChatService
	jwt          m.JWTService
	clients      map[*websocket.Conn]bool
	broadcast    chan entity.Message
	addClient    chan *websocket.Conn
	removeClient chan *websocket.Conn
	mutex        sync.Mutex
	SenderID     string
}

func NewChatHandler(service s.ChatService, jwtS m.JWTService) DomainHandler {
	return &chatHandler{
		serviceChat:  service,
		jwt:          jwtS,
		clients:      make(map[*websocket.Conn]bool),
		broadcast:    make(chan entity.Message),
		addClient:    make(chan *websocket.Conn),
		removeClient: make(chan *websocket.Conn),
		mutex:        sync.Mutex{},
	}
}
