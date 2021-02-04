package ports

import "github.com/tamercuba/ports-and-adapters-minesweeper/internal/core/domain"

// GamesRepository is the port of the MinesWeeper repository
type GamesRepository interface {
	Get(id string) (domain.Game, error)
	Save(domain.Game) error
}
