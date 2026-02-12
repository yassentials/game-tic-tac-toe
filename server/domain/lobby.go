package domain

type Lobby interface {
	AddGame(r Game) error
	FindGameByCode(code string) (Game, error)
	FindRandomGameAvailable() (Game, error)
	DeleteGameByCode(code string)

	// only shown the games that aren't full
	ListAvailableGames() []Game

	// show all games including the full ones
	ListPublicGames() []Game
}
