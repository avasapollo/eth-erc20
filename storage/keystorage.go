package storage

import (
	"context"
	"os"

	"github.com/avasapollo/eth-erc20/wallet"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
)

type KeyStorage struct {
	storage *keystore.KeyStore
}

func NewKeyStorage(keyDir string) (*KeyStorage, error) {
	ks := keystore.NewKeyStore(keyDir, keystore.StandardScryptN, keystore.StandardScryptP)
	return &KeyStorage{
		storage: ks,
	}, nil
}

func (s *KeyStorage) CreateAccount(ctx context.Context, password string) (*accounts.Account, error) {
	acc, err := s.storage.NewAccount(password)
	if err != nil {
		return nil, err
	}
	return &acc, err
}

func (s *KeyStorage) GetAccount(ctx context.Context, addr string) (*accounts.Account, error) {
	query := accounts.Account{
		Address: common.HexToAddress(addr),
	}
	add, err := s.storage.Find(query)
	if err != nil {
		if err.Error() == "no key for given address or file" {
			return nil, wallet.ErrNotFound
		}
		return nil, err
	}
	return &add, nil
}

func (s *KeyStorage) GetKey(ctx context.Context, addr, password string) (*keystore.Key, error) {
	acc, err := s.GetAccount(ctx, addr)
	if err != nil {
		return nil, err
	}

	b, err := os.ReadFile(acc.URL.Path)
	if err != nil {
		return nil, err
	}
	return keystore.DecryptKey(b, password)
}

func (s *KeyStorage) ImportJson(ctx context.Context, keyJson []byte, passphrase, newPassphrase string) (*accounts.Account, error) {
	acc, err := s.storage.Import(keyJson, passphrase, newPassphrase)
	if err != nil {
		return nil, err
	}
	return &acc, nil
}
