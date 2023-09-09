package websocket_hub

import (
	"api-fiber-mongo/service"
	"fmt"
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	"log"
	"time"
)

func GenerateCode(c *websocket.Conn) {
	userId := c.Params("userid")

	userInfo := service.GetAccountById(userId)

	code := uuid.New()

	// seria melhor utilizar o redis aqui, no momento mongo quebra o galho.
	service.TemporaryCode(code, userInfo)

	c.Conn.WriteJSON(map[string]string{"code": code.String()})
	go keepAlive(c)
	readDisconnect(c)
}

func keepAlive(c *websocket.Conn) {
	ticker := time.NewTicker(time.Second * 6)
	defer func() {
		ticker.Stop()
		c.Close()
	}()
	for {
		select {
		case <-ticker.C:
			err := c.WriteControl(websocket.PingMessage, nil, time.Now().Add(time.Second*30))
			if err != nil {
				log.Print(fmt.Sprintf("Ping error: %+v", err))
				return
			}
		}
	}

}
