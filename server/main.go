package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/yassentials/game-tic-tac-toe/server/command"
	"github.com/yassentials/game-tic-tac-toe/server/common"
	"github.com/yassentials/game-tic-tac-toe/server/infra"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		s := r.Header.Get("origin")
		fmt.Printf("origin: %s\n", s)

		return true
	},
}

const GAME_CODE_LENGTH = 5

func main() {
	lobby := infra.NewInMemoryLobby()

	handler := NewWebsocketHandler(WebsocketHandlerCommand{
		CreateGame: command.NewCreateGame(lobby, GAME_CODE_LENGTH, func() string {
			return common.GenRandomCode(GAME_CODE_LENGTH)
		}),
	})

	http.HandleFunc("/www", handler.Handle)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("failed to listen http server", err)
		return
	}
}
