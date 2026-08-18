package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/up9t/game-tic-tac-toe/backend/adapter"
	"github.com/up9t/game-tic-tac-toe/backend/command"
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
	host := flag.String("host", "localhost", "")
	port := flag.Int("port", 8080, "")

	flag.Usage = func() {
		flag.PrintDefaults()
	}

	flag.Parse()

	addr := fmt.Sprintf("%s:%d", *host, *port)

	lobby := adapter.NewInMemoryLobby()
	codeGen := adapter.SimpleRandomCode{
		Length: GAME_CODE_LENGTH,
	}

	handler := NewWebsocketHandler(WebsocketHandlerCommand{
		CreateGame: command.NewCreateGameHandler(lobby, codeGen.Generate),
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/connect", handler.Handle)

	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	finish := make(chan struct{}, 1)

	go func() {
		defer func() {
			finish <- struct{}{}
		}()

		if err := server.ListenAndServe(); err != nil {
			log.Println("failed to listen http server", err)
			return
		}
	}()

	log.Printf("Http server listening on %s \n", addr)

	<-finish
}
