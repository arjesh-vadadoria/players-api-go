package main

import (
	"PlayersApi/playerServer"
	"log"
	"net/http"
)

func main() {
	server := &playerServer.PlayerServer{Store: playerServer.NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}
