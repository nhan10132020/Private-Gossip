package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nhan10132020/chatapp/server/internal/user"
	"github.com/nhan10132020/chatapp/server/internal/ws"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
	r = gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.POST("/logout", userHandler.Logout)

	r.POST("/ws/create/room", wsHandler.CreateRoom)
	r.GET("/ws/join/room/:roomid", wsHandler.JoinRoom)
	r.GET("/ws/get/room", wsHandler.GetRooms)
	r.GET("/ws/get/client/:roomid", wsHandler.GetClients)
}

func Start(addr string) error {
	return r.Run(addr)
}