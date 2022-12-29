package processor

import (
	"context"
	"math/big"

	"github.com/avasapollo/eth-erc20/gen/go/erc20"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const gasLimit = 3000000

type Wallet interface {
	GetAccount(ctx context.Context, addr string) (*accounts.Account, error)
	GetKey(ctx context.Context, addr, password string) (*keystore.Key, error)
}

type Processor struct {
	ctrAddr       common.Address
	ownerAddr     common.Address
	ownerPassword string
	client        *ethclient.Client
	ctrClient     *erc20.Erc20
	chanID        *big.Int
	wallet        Wallet
}

func New(ctrAddr, ownerAddress common.Address, ownerPassword string, cl *ethclient.Client, wallet Wallet) (*Processor, error) {
	ctrClient, err := erc20.NewErc20(ctrAddr, cl)
	if err != nil {
		return nil, err
	}

	chanID, err := cl.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	return &Processor{
		ctrAddr:       ctrAddr,
		ownerAddr:     ownerAddress,
		ownerPassword: ownerPassword,
		client:        cl,
		ctrClient:     ctrClient,
		chanID:        chanID,
		wallet:        wallet,
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
	transactorOpt.GasLimit = gasLimit
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

type GetBalanceResult struct {
	Amount int64
}

func (p *Processor) GetBalance(ctx context.Context, address string) (*GetBalanceResult, error) {
	opt := &bind.CallOpts{}

	res, err := p.ctrClient.BalanceOf(opt, common.HexToAddress(address))
	if err != nil {
		return nil, err
	}
	return &GetBalanceResult{
		Amount: res.Int64(),
	}, nil
}

type ApproveParams struct {
	OwnerAddr     string
	OwnerPassword string
	DelegateAddr  string
	Amount        *big.Int
}

type ApproveResult struct {
	TransactionHex string
}

func (p *Processor) Approve(ctx context.Context, req *ApproveParams) (*ApproveResult, error) {
	ownerKey, err := p.wallet.GetKey(ctx, p.ownerAddr.Hex(), p.ownerPassword)
	if err != nil {
		return nil, err
	}

	delegateAddr := common.HexToAddress(req.DelegateAddr)

	gasPrice, err := p.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	transactorOpt, err := bind.NewKeyedTransactorWithChainID(ownerKey.PrivateKey, p.chanID)
	if err != nil {
		return nil, err
	}
	transactorOpt.GasLimit = gasLimit
	transactorOpt.GasPrice = gasPrice

	tr, err := p.ctrClient.Approve(transactorOpt, delegateAddr, req.Amount)
	if err != nil {
		return nil, err
	}
	return &ApproveResult{
		TransactionHex: tr.Hash().Hex(),
	}, nil
}

type AllowanceParams struct {
	OwnerAddr    string
	DelegateAddr string
}

type AllowanceResult struct {
	OwnerAddr    string
	DelegateAddr string
	Amount       int64
}

func (p *Processor) Allowance(ctx context.Context, req *AllowanceParams) (*AllowanceResult, error) {
	opt := &bind.CallOpts{}
	res, err := p.ctrClient.Allowance(opt, common.HexToAddress(req.OwnerAddr), common.HexToAddress(req.DelegateAddr))
	if err != nil {
		return nil, err
	}
	return &AllowanceResult{
		OwnerAddr:    req.OwnerAddr,
		DelegateAddr: req.DelegateAddr,
		Amount:       res.Int64(),
	}, nil
}

type TransferFromParams struct {
	DelegateAddr     string
	DelegatePassword string
	OwnerAddr        string
	BuyerAddr        string
	Amount           *big.Int
}

type TransferFromResult struct {
	TransactionHex string
	DelegateAddr   string
	OwnerAddr      string
	BuyerAddr      string
}

func (p *Processor) TransferFrom(ctx context.Context, req *TransferFromParams) (*TransferFromResult, error) {
	delegteKey, err := p.wallet.GetKey(ctx, req.DelegateAddr, req.DelegatePassword)
	if err != nil {
		return nil, err
	}

	ownerAddr := common.HexToAddress(req.OwnerAddr)
	buyerAddr := common.HexToAddress(req.BuyerAddr)

	gasPrice, err := p.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	transactorOpt, err := bind.NewKeyedTransactorWithChainID(delegteKey.PrivateKey, p.chanID)
	if err != nil {
		return nil, err
	}
	transactorOpt.GasLimit = gasLimit
	transactorOpt.GasPrice = gasPrice

	tr, err := p.ctrClient.TransferFrom(transactorOpt, ownerAddr, buyerAddr, req.Amount)
	if err != nil {
		return nil, err
	}
	return &TransferFromResult{
		TransactionHex: tr.Hash().Hex(),
		DelegateAddr:   req.DelegateAddr,
		OwnerAddr:      req.OwnerAddr,
		BuyerAddr:      req.BuyerAddr,
	}, nil
}
