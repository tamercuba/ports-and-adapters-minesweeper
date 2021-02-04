package usecases

import (
	"errors"

	"github.com/tamercuba/ports-and-adapters-minesweeper/internal/core/domain"
	"github.com/tamercuba/ports-and-adapters-minesweeper/internal/core/ports"
	"github.com/tamercuba/ports-and-adapters-minesweeper/pkg/uidgen"
)

// Service is oyut entry point to the core app
type Service struct {
	gamesRepository ports.GamesRepository
	uidGen          uidgen.UIDGen
}

// New is the Service constructor
func New(repository ports.GamesRepository) *Service {
	return &Service{
		gamesRepository: repository,
	}
}

// Get returns a saved game by id
func (service *Service) Get(id string) (domain.Game, error) {
	game, err := service.gamesRepository.Get(id)
	if err != nil {
		return domain.Game{}, errors.New("Something failed when getting game by id")
	}
	return game, nil
}

// Create creates a new game
func (service *Service) Create(name string, size uint, bombs uint) (domain.Game, error) {
	if bombs >= size*size {
		return domain.Game{}, errors.New("The borad must be larger than bombs number")
	}
	game := domain.NewGame(service.uidGen.New(), name, size, bombs)

	if err := service.gamesRepository.Save(game)
}
