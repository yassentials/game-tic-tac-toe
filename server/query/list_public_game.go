package query

import (
	"log"

	"github.com/yassentials/game-tic-tac-toe/server/domain"
)

// filter by
// accessibility [all,joinable]
type ListFilterFlag int
type ListFilterAccessibilityFlag ListFilterFlag

// type OtherFilterExample ListFilterFlag

const (
	LIST_FILTER_NONE ListFilterAccessibilityFlag = 1 << iota
	LIST_FILTER_ACCESSIBILITY_JOINABLE
	LIST_FILTER_ACCESSIBILITY_ALL
)

type ListPublicGameHandler struct {
	lobby domain.Lobby
}

func NewListPublicGameHandler(lobby domain.Lobby) *ListPublicGameHandler {
	return &ListPublicGameHandler{
		lobby: lobby,
	}
}

type ListPublicGameQuery struct {
	FilterFlags ListFilterFlag
}

func (h *ListPublicGameHandler) Handle(query ListPublicGameQuery) []domain.Game {
	if query.FilterFlags == ListFilterFlag(LIST_FILTER_NONE) {
		query.FilterFlags |= ListFilterFlag(LIST_FILTER_ACCESSIBILITY_JOINABLE)
	}

	if query.FilterFlags&ListFilterFlag(LIST_FILTER_ACCESSIBILITY_JOINABLE) > 0 {
		return h.lobby.ListAvailableGames()
	}

	if query.FilterFlags&ListFilterFlag(LIST_FILTER_ACCESSIBILITY_ALL) > 0 {
		return h.lobby.ListPublicGames()
	}

	log.Printf("[List Game] Filter unavailable: %b.\n", query.FilterFlags)

	return []domain.Game{}
}
