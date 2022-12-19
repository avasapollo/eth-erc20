package wallets

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/accounts"
)

var (
	ErrNotFound = errors.New("not found")
)

type Storage interface {
	CreateAccount(ctx context.Context, password string) (*accounts.Account, error)
}

type Service struct {
	storage Storage
}

func New(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) CreateAccount(ctx context.Context, password string) (*accounts.Account, error) {
	return s.storage.CreateAccount(ctx, password)
}
