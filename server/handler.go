package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/yassentials/game-tic-tac-toe/server/command"
	"github.com/yassentials/game-tic-tac-toe/server/domain"
	"github.com/yassentials/game-tic-tac-toe/server/event"
)

type RequestType int
type ResponseType int

const (
	TYPE_REQ_GAME_JOIN_RANDOM RequestType = iota
	TYPE_REQ_GAME_JOIN_CODE
	TYPE_REQ_GAME_CREATE

	TYPE_REQ_ACTION_TAKE
)

const (
	TYPE_RES_GAME_FULL ResponseType = iota

	TYPE_RES_ACTION_TAKE_INVALID
)

type BaseMessage[T RequestType | ResponseType] struct {
	Type T `json:"typ"`
}

type CreateGameRequest struct {
	BaseMessage[RequestType]

	Name       string           `json:"nam"`
	Character  domain.Character `json:"cha"`
	Visibility domain.GameType  `json:"vis"`
}

type JoinGameWithCodeRequest struct {
	BaseMessage[RequestType]

	Name      string           `json:"nam"`
	Character domain.Character `json:"cha"`
	Code      string           `json:"cod"`
}

type JoinRandomGameRequest struct {
	BaseMessage[RequestType]

	Name      string           `json:"nam"`
	Character domain.Character `json:"cha"`
}

type TakePositionRequest struct {
	BaseMessage[RequestType]

	Index int `json:"ind"`
}

type WebsocketHandlerCommand struct {
	CreateGame     command.CreateGameHandler
	JoinGameByCode command.JoinGameByCodeHandler
	JoinRandomGame command.JoinRandomGameHandler
	LeaveGame      command.LeaveGameHandler
	TakePosition   command.TakePositionHandler
}

type WebsocketHandler struct {
	cmd WebsocketHandlerCommand
}

func NewWebsocketHandler(cmd WebsocketHandlerCommand) WebsocketHandler {
	return WebsocketHandler{
		cmd,
	}
}

func (h WebsocketHandler) Handle(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err != nil {
		log.Printf("[Connection Upgrade] failed: %s\n", err.Error())
		return
	}

	defer conn.Close()

	var (
		activeGame domain.Game
		player     domain.Player
		ready      = make(chan struct{})
	)

	go func() {
		<-ready

		unlistener := activeGame.GetEventManager().Listen(event.EVENT_TAKE_POSITION_SUCCEED, func(e domain.Event[any]) {
			data, ok := e.GetData().(event.TakePositionSucceedEventData)

			if !ok {
				log.Println("[Interface conversion] failed")
				return
			}

			conn.WriteJSON(data)
		})

		<-ctx.Done()

		h.cmd.LeaveGame.Handle(command.LeaveGameCommand{
			Player: player,
			Game:   activeGame,
		})

		unlistener()
	}()

	for {
		var baseReq BaseMessage[RequestType]

		msgType, msg, err := conn.ReadMessage()

		log.Printf("Message type: %s\n", msgType)

		if err != nil {
			log.Printf("[Read Message] failed: %s\n", err.Error())
			continue
		}

		if err = json.Unmarshal(msg, &baseReq); err != nil {
			log.Printf("[Parse JSON] failed: %s\n", err.Error())
			continue
		}

		switch baseReq.Type {
		case TYPE_REQ_GAME_CREATE:
			var req CreateGameRequest

			if err = json.Unmarshal(msg, &req); err != nil {
				log.Printf("[Parse JSON] failed: %s\n", err.Error())
				continue
			}

			activeGame, player, err = h.cmd.CreateGame.Handle(command.CreateGameCommand{
				Type:             domain.GAME_TYPE_PUBLIC,
				PlayerName:       req.Name,
				PlayerCharacater: req.Character,
			})

			if err != nil {
				log.Printf("[Create Game] failed: %s\n", err.Error())
				continue
			}

			close(ready)
		case TYPE_REQ_GAME_JOIN_RANDOM:
			var req JoinRandomGameRequest

			if err = json.Unmarshal(msg, &req); err != nil {
				log.Printf("[Parse JSON] failed: %s\n", err.Error())
				continue
			}

			activeGame, player, err = h.cmd.JoinRandomGame.Handle(command.JoinRandomGameCommand{
				PlayerName:       req.Name,
				PlayerCharacater: req.Character,
			})

			if err != nil {
				log.Printf("[Join Random] failed: %s\n", err.Error())
				continue
			}

			close(ready)
		case TYPE_REQ_GAME_JOIN_CODE:
			var req JoinGameWithCodeRequest

			if err = json.Unmarshal(msg, &req); err != nil {
				log.Printf("[Parse JSON] failed: %s\n", err.Error())
				continue
			}

			activeGame, player, err = h.cmd.JoinGameByCode.Handle(command.JoinGameByCodeCommand{
				PlayerName:       req.Name,
				PlayerCharacater: req.Character,
				Code:             req.Code,
			})

			if err != nil {
				log.Printf("[Join Code] failed: %s\n", err.Error())
				continue
			}

			close(ready)
		case TYPE_REQ_ACTION_TAKE:
			var req TakePositionRequest

			if err = json.Unmarshal(msg, &req); err != nil {
				log.Printf("[Parse JSON] failed: %s\n", err.Error())
				continue
			}

			if err := h.cmd.TakePosition.Handle(command.TakePositionCommand{
				Game:   activeGame,
				Player: player,
				Index:  req.Index,
			}); err != nil {
				log.Printf("[Action Take] failed: %s\n", err.Error())
				continue
			}
		default:
			continue
		}
	}
}
