package ports

import "github.com/tamercuba/ports-and-adapters-minesweeper/internal/core/domain"

// GamesService is the port of the Game service MinesWeeper
type GamesService interface {
	Get(id string) (domain.Game, error)
	Create(name string, size uint, bombs uint) (domain.Game, error)
	Reveal(id string, row uint, col uint) (domain.Game, error)
}
