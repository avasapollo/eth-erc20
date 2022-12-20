package handlers

import (
	"context"

	v1 "github.com/avasapollo/eth-erc20/gen/go/wallet/v1"
	"github.com/avasapollo/eth-erc20/gen/go/wallet/v1/walletv1connect"
	connect_go "github.com/bufbuild/connect-go"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

var _ walletv1connect.WalletServiceHandler = (*WalletHandler)(nil)

type WalletService interface {
	CreateAccount(ctx context.Context, password string) (*accounts.Account, error)
	GetAccount(ctx context.Context, addr string) (*accounts.Account, error)
	ImportJson(ctx context.Context, keyJson []byte, passphrase, newPassphrase string) (*accounts.Account, error)
	GetKey(ctx context.Context, addr, password string) (*keystore.Key, error)
}

type WalletHandler struct {
	wallet WalletService
}

func NewWalletHandler(walletSvc WalletService) *WalletHandler {
	return &WalletHandler{
		wallet: walletSvc,
	}
}

func toCreateAccountResp(acc *accounts.Account) *v1.CreateAccountResp {
	return &v1.CreateAccountResp{
		Address: acc.Address.Hex(),
	}
}

func (s *WalletHandler) CreateAccount(ctx context.Context, req *connect_go.Request[v1.CreateAccountReq]) (*connect_go.Response[v1.CreateAccountResp], error) {
	acc, err := s.wallet.CreateAccount(ctx, req.Msg.GetPassword())
	if err != nil {
		return nil, err
	}
	return connect_go.NewResponse(toCreateAccountResp(acc)), nil
}

func toGetAccountResp(acc *accounts.Account) *v1.GetAccountResp {
	return &v1.GetAccountResp{
		Address: acc.Address.Hex(),
	}
}

func (s *WalletHandler) GetAccount(ctx context.Context, req *connect_go.Request[v1.GetAccountReq]) (*connect_go.Response[v1.GetAccountResp], error) {
	acc, err := s.wallet.GetAccount(ctx, req.Msg.GetAddress())
	if err != nil {
		return nil, err
	}
	return connect_go.NewResponse(toGetAccountResp(acc)), nil
}

func toImportAccountResp(acc *accounts.Account) *v1.ImportAccountResp {
	return &v1.ImportAccountResp{
		Address: acc.Address.Hex(),
	}
}

func (s *WalletHandler) ImportAccount(ctx context.Context, req *connect_go.Request[v1.ImportAccountReq]) (*connect_go.Response[v1.ImportAccountResp], error) {
	acc, err := s.wallet.ImportJson(ctx, req.Msg.GetKeyJson(), req.Msg.GetPassphrase(), req.Msg.GetNewPassphrase())
	if err != nil {
		return nil, err
	}
	return connect_go.NewResponse(toImportAccountResp(acc)), nil
}
