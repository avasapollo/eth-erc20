package main

import (
	"context"

	"github.com/avasapollo/eth-erc20/storage"
	"github.com/avasapollo/eth-erc20/wallet"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type config struct {
	OwnerPassword    string `envconfig:"OWNER_PASSWORD"`
	ReceiverPassword string `envconfig:"RECEIVER_PASSWORD"`
	DelegatePassword string `envconfig:"DELEGATE_PASSWORD"`
	BuyerPassword    string `envconfig:"BUYER_PASSWORD"`
	KeyDir           string `envconfig:"KEY_DIR"`
}

func main() {
	ctx := context.Background()
	lgr := logrus.New().WithFields(logrus.Fields{
		"app": "eoa-creator",
	})

	c := new(config)
	err := envconfig.Process("", c)
	if err != nil {
		lgr.WithError(err).Fatal("can't load config")
	}

	keyStorage, err := storage.NewKeyStorage(c.KeyDir)
	if err != nil {
		lgr.WithError(err).Fatal("can't create the key storage")
	}

	walletSvc := wallet.New(keyStorage)

	// create contract owner wallet
	owner, err := walletSvc.CreateAccount(ctx, c.OwnerPassword)
	if err != nil {
		lgr.WithError(err).Fatal("can't create the owner account")
	}

	lgr.Infof("owner address: %s", owner.Address.Hex())

	// create receiver wallet
	receiver, err := walletSvc.CreateAccount(ctx, c.OwnerPassword)
	if err != nil {
		lgr.WithError(err).Fatal("can't create the receiver account")
	}

	lgr.Infof("receiver address: %s", receiver.Address.Hex())

	// create delegate wallet
	delegate, err := walletSvc.CreateAccount(ctx, c.DelegatePassword)
	if err != nil {
		lgr.WithError(err).Fatal("can't create the delegate account")
	}

	lgr.Infof("delegate address: %s", delegate.Address.Hex())

	// create buyer wallet
	buyer, err := walletSvc.CreateAccount(ctx, c.BuyerPassword)
	if err != nil {
		lgr.WithError(err).Fatal("can't create the buyer account")
	}

	lgr.Infof("buyer address: %s", buyer.Address.Hex())
}
