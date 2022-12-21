package main

import (
	"context"

	"github.com/avasapollo/eth-erc20/storage"
	"github.com/avasapollo/eth-erc20/wallet"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type config struct {
	OwnerPassword string `envconfig:"OWNER_PASSWORD"`
	OwnerBalance  int64  `envconfig:"OWNER_BALANCE"`
	KeyDir        string `envconfig:"KEY_DIR"`
}

func main() {
	ctx := context.Background()
	lgr := logrus.New().WithFields(logrus.Fields{
		"app": "owner-creator",
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
}
