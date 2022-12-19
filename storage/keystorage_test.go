package storage

import (
	"context"
	"errors"
	"testing"

	"github.com/avasapollo/eth-erc20/wallet"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/stretchr/testify/require"
)

type testSuite struct {
	storage *keystore.KeyStore
}

func newTestSuite() *testSuite {
	path := "./tmp"
	return &testSuite{
		storage: keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP),
	}
}

func TestKeyStorage_CreateAccount(t *testing.T) {
	type args struct {
		ctx      context.Context
		password string
	}
	tests := []struct {
		name string
		args args
		want func(got *accounts.Account, err error)
	}{
		{
			name: "ok store account",
			args: args{
				ctx:      context.Background(),
				password: "password",
			},
			want: func(got *accounts.Account, err error) {
				require.NoError(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			suite := newTestSuite()
			s := &KeyStorage{
				storage: suite.storage,
			}
			got, err := s.CreateAccount(tt.args.ctx, tt.args.password)
			tt.want(got, err)
		})
	}
}

func TestKeyStorage_GetAccount(t *testing.T) {
	t.Run("existing account ", func(t *testing.T) {
		suite := newTestSuite()
		ctx := context.Background()
		password := "password"
		s := &KeyStorage{
			storage: suite.storage,
		}

		// create account
		acc, err := s.CreateAccount(ctx, password)
		if err != nil {
			t.Fatal(err)
		}

		got, err := s.GetAccount(ctx, acc.Address.Hex())
		if err != nil {
			t.Fatal(err)
		}
		t.Log(got.Address.Hex())
	})

	t.Run("error not found ", func(t *testing.T) {
		suite := newTestSuite()
		ctx := context.Background()
		s := &KeyStorage{
			storage: suite.storage,
		}

		// 0xa1A5B87fEe23dD5ac770cB2E8D6Cf7665d0638c1 not exist
		_, err := s.GetAccount(ctx, "0xa1A5B87fEe23dD5ac770cB2E8D6Cf7665d0638c1")
		require.True(t, errors.Is(err, wallet.ErrNotFound))
	})
}

func TestKeyStorage_GetKey(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		suite := newTestSuite()
		ctx := context.Background()
		password := "password"

		s := &KeyStorage{
			storage: suite.storage,
		}

		acc, err := s.CreateAccount(ctx, password)
		if err != nil {
			t.Fatal(err)
		}

		key, err := s.GetKey(ctx, acc.Address.Hex(), password)
		if err != nil {
			t.Fatal(err)
		}

		require.NotEmpty(t, key.Id.String())
	})
}
