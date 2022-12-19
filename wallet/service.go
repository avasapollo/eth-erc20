package wallet

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

var (
	ErrNotFound = errors.New("not found")
)

type Storage interface {
	CreateAccount(ctx context.Context, password string) (*accounts.Account, error)
	GetAccount(ctx context.Context, addr string) (*accounts.Account, error)
	ImportJson(ctx context.Context, keyJson []byte, passphrase, newPassphrase string) (*accounts.Account, error)
	GetKey(ctx context.Context, addr, password string) (*keystore.Key, error)
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

func (s *Service) GetAccount(ctx context.Context, addr string) (*accounts.Account, error) {
	return s.storage.GetAccount(ctx, addr)
}

func (s *Service) ImportJson(ctx context.Context, keyJson []byte, passphrase, newPassphrase string) (*accounts.Account, error) {
	return s.storage.ImportJson(ctx, keyJson, passphrase, newPassphrase)
}

func (s *Service) GetKey(ctx context.Context, addr, password string) (*keystore.Key, error) {
	return s.storage.GetKey(ctx, addr, password)
}
