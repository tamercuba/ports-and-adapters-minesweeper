// .internal/core/domain/domain.go

package domain

const (
	// GameStateWon saves the Won state
	GameStateWon = "WON"
	// GameStateLost saves the Lost state
	GameStateLost = "LOST"
	// GameStateNew saves the New state
	GameStateNew = "NEW"
)

// Game is our main entity
type Game struct {
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	State         string        `json:"state"`
	BoardSettings BoardSettings `json:"board_settings"`
	Board         Board         `json:"board"`
}

// NewGame is the Game constructor
func NewGame(id string, name string, size uint, bombs uint) Game {
	return Game{
		ID:    id,
		Name:  name,
		State: GameStateNew,
		BoardSettings: BoardSettings{
			Size:  size,
			Bombs: bombs,
		},
		Board: NewBoard(size, bombs),
	}
}

// IsOver checks if the game is over
func (game *Game) IsOver() bool {
	return game.State == GameStateLost || game.State == GameStateWon
}
