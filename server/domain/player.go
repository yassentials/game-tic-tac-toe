package domain

type Player interface {
	GetName() string
	GetCharacter() Character
	SetCharacter(char Character)
}

type GamePlayer struct {
	name      string
	character Character
}

func NewGamePlayer(name string, character Character) *GamePlayer {
	return &GamePlayer{
		name,
		character,
	}
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
