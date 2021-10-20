package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) initGameRoutes(api *gin.RouterGroup) {
	g := api.Group("/games")
	{
		g.GET("", h.saveGame)
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) saveGame(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Errorf("error set connection - %s", err.Error())
		return
	}

	// start session
	for {
		mt, m, err := ws.ReadMessage()

		if err != nil {
			log.Errorf("error read m - %s", err.Error())
			continue
		}

		log.Infof("message received - %s", string(m))

		if err = ws.WriteMessage(mt, []byte("ok")); err != nil {
			log.Errorf("error write messsage - %s", err.Error())
		}
	}
}
