package servers

import (
	"net/http"
	"time"

	"github.com/besanh/soa/common/log"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
}

func NewServer() *Server {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"service": "soa",
			"version": "1.0.0",
			"time":    time.Now().Unix(),
		})
	})
	server := &Server{Engine: engine}
	return server
}

func (server *Server) Start(port string) {
	v := make(chan struct{})
	go func() {
		if err := server.Engine.Run(":" + port); err != nil {
			log.Errorf("failed to start service")
			close(v)
		}
	}()
	log.Debugf("service is listening on port %v", port)
	<-v
}
