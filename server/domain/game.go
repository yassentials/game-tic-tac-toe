package domain

import (
	"fmt"
	"sync"
)

type GameState int
type GameType int
type GameResult int

const (
	GAME_RESULT_NONE GameResult = iota
	GAME_RESULT_EITHER_WIN
	GAME_RESULT_DRAW
)

const (
	GAME_STATE_MATCHMAKING GameState = iota
	GAME_STATE_PLAYING
	GAME_STATE_RESULT
)

const (
	// can be discovered publicly
	GAME_TYPE_PUBLIC GameType = iota

	// use code to join
	GAME_TYPE_PRIVATE
)

type Room interface {
	// unique code for joining the room
	GetCode() string
	IsFull() bool
	Join(player Player) error
	Leave(player Player)
}

type Game interface {
	Room

	GetEventManager() EventManager[any]
	GetResult() GameResult
	GetType() GameType
	GetState() GameState
	SetState(state GameState) error
	Restart()
	GetWinner() Player
	TakePosition(player Player, index int) error
}

type BaseGame struct {
	// type is reserved keyword here, so I had to use gameType instead
	result       GameResult
	gameType     GameType
	state        GameState
	code         string
	capacity     int
	players      []Player
	winner       Player
	board        []Character
	mu           sync.RWMutex
	turn         int
	eventManager EventManager[any]
}

func NewBaseGame(capacity int, gameType GameType, eventManager EventManager[any], codeGen func() string) *BaseGame {
	game := &BaseGame{
		turn:         0,
		gameType:     gameType,
		state:        GAME_STATE_MATCHMAKING,
		code:         codeGen(),
		capacity:     capacity,
		players:      make([]Player, capacity),
		board:        make([]Character, 3*3),
		eventManager: eventManager,
	}

	return game
}

func (g *BaseGame) Restart() {
	g.turn = 0
	g.state = GAME_STATE_PLAYING
	g.board = make([]Character, 3*3)
}

func (g *BaseGame) GetType() GameType {
	return g.gameType
}

func (g *BaseGame) GetResult() GameResult {
	return g.result
}

func (g *BaseGame) GetState() GameState {
	return g.state
}

func (g *BaseGame) SetState(state GameState) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.state = state

	return nil
}

func (g *BaseGame) GetEventManager() EventManager[any] {
	return g.eventManager
}
func (g *BaseGame) TakePosition(player Player, index int) error {
	g.mu.Lock()

	if index < 0 || index >= len(g.board) {
		g.mu.Unlock()
		return fmt.Errorf("[Invalid Index] must between %d and %d\n", 0, len(g.board)-1)
	}

	if player.GetCharacter() == CHAR_NONE {
		g.mu.Unlock()
		return fmt.Errorf("[Invalid Character]")
	}

	g.board[index] = player.GetCharacter()
	g.mu.Unlock()

	g.determineResult()
	g.acquireTurn()

	return nil
}

func (g *BaseGame) GetCode() string {
	return g.code
}

func (g *BaseGame) IsFull() bool {
	g.mu.RLock()
	defer g.mu.RUnlock()

	return len(g.players) >= g.capacity
}

func (g *BaseGame) Join(player Player) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if len(g.players) >= g.capacity {
		return fmt.Errorf("[Room Full]")
	}

	g.players = append(g.players, player)

	return nil
}

func (g *BaseGame) Leave(player Player) {

}

func (g *BaseGame) getCurrentPlayer() Player {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return g.players[g.turn]
}

func (g *BaseGame) GetWinner() Player {
	return g.winner
}

// should be called before acquireTurn()
func (g *BaseGame) determineResult() {
	g.mu.Lock()
	defer g.mu.Unlock()

	WINNING_BOARD := []int{
		// horizontal line
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,

		// vertical line
		0, 3, 6,
		1, 4, 7,
		2, 5, 8,

		// "/" line
		0, 4, 8,

		// "\" line
		2, 4, 6,
	}

	for i := 0; i < len(WINNING_BOARD); i += 3 {
		bd := WINNING_BOARD

		a, b, c := bd[i], bd[i+1], bd[i+2]

		if a == b && b == c {
			g.result = GAME_RESULT_EITHER_WIN
			g.state = GAME_STATE_RESULT
			g.winner = g.getCurrentPlayer()
			return
		}
	}

	for _, b := range g.board {
		if b != CHAR_NONE {
			g.result = GAME_RESULT_NONE
			return
		}
	}

	g.state = GAME_STATE_RESULT
	g.result = GAME_RESULT_DRAW
}

// should be called after determineResult()
func (g *BaseGame) acquireTurn() {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.turn %= g.capacity
}
