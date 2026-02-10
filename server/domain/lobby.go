package domain

type Lobby interface {
	AddGame(r Game) error
	FindGameByCode(code string) (Game, error)
	ListAvailableGames() []Game
	FindRandomGameAvailable() Game
	DeleteGameByCode(code string)
}
