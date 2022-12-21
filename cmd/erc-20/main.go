package main

import (
	"context"
	"net"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/avasapollo/eth-erc20/gen/go/balance/v1/balancev1connect"
	"github.com/avasapollo/eth-erc20/gen/go/transfer/v1/transferv1connect"
	"github.com/avasapollo/eth-erc20/gen/go/wallet/v1/walletv1connect"
	"github.com/avasapollo/eth-erc20/handlers"
	"github.com/avasapollo/eth-erc20/processor"
	"github.com/avasapollo/eth-erc20/server"
	"github.com/avasapollo/eth-erc20/storage"
	"github.com/avasapollo/eth-erc20/wallet"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type config struct {
	NetworkURL      string `envconfig:"NETWORK_URL"`
	OwnerAddress    string `envconfig:"OWNER_ADDRESS"`
	OwnerPassword   string `envconfig:"OWNER_PASSWORD"`
	ContractAddress string `envconfig:"CONTRACT_ADDRESS"`
	KeyDir          string `envconfig:"KEY_DIR"`
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	lgr := logrus.New().WithFields(logrus.Fields{
		"app": "erc-20",
	})

	c := new(config)
	err := envconfig.Process("", c)
	if err != nil {
		lgr.WithError(err).Fatal("can't load the config")
	}

	keyStorage, err := storage.NewKeyStorage(c.KeyDir)
	if err != nil {
		lgr.WithError(err).Fatal("can't create the key storage")
	}

	walletSvc := wallet.New(keyStorage)

	walletHandler := handlers.NewWalletHandler(walletSvc)

	client, err := ethclient.DialContext(ctx, c.NetworkURL)
	if err != nil {
		lgr.WithError(err).Fatal("can't create the client")
	}

	processorSvc, err := processor.New(
		common.HexToAddress(c.ContractAddress),
		common.HexToAddress(c.OwnerAddress),
		c.OwnerPassword,
		client,
		walletSvc,
	)
	if err != nil {
		lgr.WithError(err).Fatal("can't create the processor")
	}
	transferHandler := handlers.NewTransferHandler(processorSvc)
	balanceHandler := handlers.NewBalanceHandler(processorSvc)

	mux := http.NewServeMux()
	mux.Handle(walletv1connect.NewWalletServiceHandler(walletHandler))
	mux.Handle(transferv1connect.NewTransferServiceHandler(transferHandler))
	mux.Handle(balancev1connect.NewBalanceServiceHandler(balanceHandler))

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		lgr.WithError(err).Fatal("failed to create listener")
	}
	srv := server.NewServer(mux, lis)
	go func() {
		if err := srv.Start(); err != nil {
			lgr.Fatal(err, "failed to start server")
		}
	}()
	lgr.Info("application started", "addr", srv.Addr())

	// wait for shutdown
	<-ctx.Done()

	lgr.Info("application shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		lgr.Error(err, "failed to shutdown server")
	}
}
