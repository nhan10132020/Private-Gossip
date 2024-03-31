package main

import (
	"log"

	"github.com/nhan10132020/chatapp/server/db"
	"github.com/nhan10132020/chatapp/server/internal/user"
	"github.com/nhan10132020/chatapp/server/internal/ws"
	"github.com/nhan10132020/chatapp/server/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatal("COULD NOT INIT DATABASE CONN", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	router.Start("localhost:8080")
}
