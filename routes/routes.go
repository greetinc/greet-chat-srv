package routes

import (
	"greet-chat-srv/configs"

	"github.com/greetinc/greet-middlewares/middlewares"
	"github.com/labstack/echo/v4"

	h_chat "greet-chat-srv/handlers/chat"
	r_chat "greet-chat-srv/repositories/chat"
	s_chat "greet-chat-srv/services/chat"
)

var (
	DB = configs.InitDB()

	JWT                 = middlewares.NewJWTService()
	chatR               = r_chat.NewChatRepository(DB)
	chatS               = s_chat.NewChatService(chatR, JWT)
	chatH               = h_chat.NewChatHandler(chatS, JWT)
	chatHandlerInstance = h_chat.NewChatHandler(chatS, JWT)
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/history", chatH.GetMessagesBetweenUsers)
	e.GET("/ws", func(c echo.Context) error {
		chatHandlerInstance.HandleConnections(c.Response(), c.Request())
		return nil
	})

	// Goroutine untuk menangani pesan pada koneksi WebSocket
	go chatHandlerInstance.HandleMessages()

	return e
}
