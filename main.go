package main

import (
	"github.com/jorgeme09/twittor/db"
	"github.com/jorgeme09/twittor/handlers"
	"log"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Sin conexion con DB")
		return
	}
	handlers.Handlers()
}
