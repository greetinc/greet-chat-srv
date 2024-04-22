package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (u *chatHandler) GetMessagesBetweenUsers(c echo.Context) error {

	sender_id := c.QueryParam("sender_id")
	receiver_id := c.QueryParam("receiver_id")

	messages, err := u.GetMessagesFromDatabase(sender_id, receiver_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch messages")
	}

	return c.JSON(http.StatusOK, messages)
}
