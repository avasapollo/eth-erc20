package processor

import (
	"context"
	"math/big"

	erc20 "github.com/avasapollo/eth-erc20/gen"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Wallet interface {
	GetAccount(ctx context.Context, addr string) (*accounts.Account, error)
	GetKey(ctx context.Context, addr, password string) (*keystore.Key, error)
}

type Processor struct {
	ctrAddr   common.Address
	client    *ethclient.Client
	ctrClient *erc20.Erc20
	chanID    *big.Int
	wallet    Wallet
}

func New(ctrAddr common.Address, cl *ethclient.Client, wallet Wallet) (*Processor, error) {
	ctrClient, err := erc20.NewErc20(ctrAddr, cl)
	if err != nil {
		return nil, err
	}

	chanID, err := cl.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	return &Processor{
		ctrAddr:   ctrAddr,
		client:    cl,
		ctrClient: ctrClient,
		chanID:    chanID,
		wallet:    wallet,
	}, nil
}

type TransferParams struct {
	SenderAddr     string
	SenderPassword string
	ReceiverAddr   string
	Amount         *big.Int
}

type TransferResult struct {
	TransactionHex string
	SenderAddr     string
	ReceiverAddr   string
}

func (p *Processor) Transfer(ctx context.Context, req *TransferParams) (*TransferResult, error) {
	senderKey, err := p.wallet.GetKey(ctx, req.SenderAddr, req.SenderPassword)
	if err != nil {
		return nil, err
	}

	receiverAddr := common.HexToAddress(req.ReceiverAddr)

	gasPrice, err := p.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	transactorOpt, err := bind.NewKeyedTransactorWithChainID(senderKey.PrivateKey, p.chanID)
	if err != nil {
		return nil, err
	}
	transactorOpt.GasLimit = 3000000
	transactorOpt.GasPrice = gasPrice

	tr, err := p.ctrClient.Transfer(transactorOpt, receiverAddr, req.Amount)
	if err != nil {
		return nil, err
	}
	return &TransferResult{
		TransactionHex: tr.Hash().Hex(),
		SenderAddr:     req.SenderAddr,
		ReceiverAddr:   req.ReceiverAddr,
	}, nil
}
