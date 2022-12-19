package main

import (
	"context"
	"math/big"

	"github.com/avasapollo/eth-erc20/converter"
	erc20 "github.com/avasapollo/eth-erc20/gen"
	"github.com/avasapollo/eth-erc20/storage"
	"github.com/avasapollo/eth-erc20/wallet"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type config struct {
	NetworkURL    string `envconfig:"NETWORK_URL"`
	OwnerPassword string `envconfig:"OWNER_PASSWORD"`
	OwnerBalance  int64  `envconfig:"OWNER_BALANCE"`
}

func main() {
	ctx := context.Background()
	lgr := logrus.New().WithFields(logrus.Fields{
		"app": "erc20-playground",
	})

	c := new(config)
	err := envconfig.Process("", c)
	if err != nil {
		lgr.WithError(err).Fatal("can't load config")
	}

	client, err := ethclient.DialContext(ctx, c.NetworkURL)
	if err != nil {
		lgr.WithError(err).Fatal("can't create the client")
	}

	keyStorage, err := storage.NewKeyStorage("./wallets")
	if err != nil {
		lgr.WithError(err).Fatal("can't create the key storage")
	}

	walletSvc := wallet.New(keyStorage)

	// create contract owner wallet
	owner, err := walletSvc.CreateAccount(ctx, c.OwnerPassword)
	if err != nil {
		lgr.WithError(err).Fatal("can't create the owner account")
	}

	ownerKey, err := walletSvc.GetKey(ctx, owner.Address.Hex(), c.OwnerPassword)
	if err != nil {
		lgr.WithError(err).Fatal("can't get owner private key")
	}

	add := crypto.PubkeyToAddress(ownerKey.PrivateKey.PublicKey)

	nonce, err := client.PendingNonceAt(ctx, add)
	if err != nil {
		lgr.WithError(err).Fatal("can't get owner nonce")
	}
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		lgr.WithError(err).Fatal("can't get suggested gas price")
	}

	chanID, err := client.NetworkID(ctx)
	if err != nil {
		lgr.WithError(err).Fatal("can't get chanID")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(ownerKey.PrivateKey, chanID)
	if err != nil {
		lgr.WithError(err).Fatal("can't create auth request to deploy contract")
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = 3000000
	auth.Nonce = big.NewInt(int64(nonce))

	conAddr, tr, _, err := erc20.DeployErc20(auth, client, big.NewInt(c.OwnerBalance))
	if err != nil {
		lgr.WithError(err).Fatal("can't create auth request to deploy contract")
	}

	trCost, _ := converter.FromWei(converter.Ether, tr.Cost())
	lgr.WithFields(logrus.Fields{
		"tr_price":         trCost,
		"tr_hash":          tr.Hash().Hex(),
		"contract_address": conAddr.Hex(),
		"chan_id":          chanID,
	}).Info("contract deployed")
}
