package domain

import "math/rand/v2"

type Player interface {
	GetId() int
	GetName() string
	GetCharacter() Character
	SetCharacter(char Character)
}

type GamePlayer struct {
	id        int
	name      string
	character Character
}

func NewGamePlayer(name string, character Character) *GamePlayer {
	return &GamePlayer{
		id:        rand.Int(),
		name:      name,
		character: character,
	}
}
func (p *GamePlayer) GetId() int {
	return p.id
}

func (p *GamePlayer) GetName() string {
	return p.name
}

func (p *GamePlayer) GetCharacter() Character {
	return p.character
}

func (p *GamePlayer) SetCharacter(char Character) {
	p.character = char
}
