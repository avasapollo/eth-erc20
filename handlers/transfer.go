package handlers

import (
	"context"
	"math/big"

	v1 "github.com/avasapollo/eth-erc20/gen/go/transfer/v1"
	"github.com/avasapollo/eth-erc20/gen/go/transfer/v1/transferv1connect"
	"github.com/avasapollo/eth-erc20/processor"
	connect_go "github.com/bufbuild/connect-go"
)

var _ transferv1connect.TransferServiceHandler = (*TransferHandler)(nil)

type ProcessorService interface {
	Transfer(ctx context.Context, req *processor.TransferParams) (*processor.TransferResult, error)
}

type TransferHandler struct {
	processor ProcessorService
}

func NewTransferHandler(processorSvc ProcessorService) *TransferHandler {
	return &TransferHandler{
		processor: processorSvc,
	}
}

func toTransferParams(req *v1.TransferReq) *processor.TransferParams {
	return &processor.TransferParams{
		SenderAddr:     req.GetSenderAddress(),
		SenderPassword: req.GetSenderPassword(),
		ReceiverAddr:   req.GetReceiverAddress(),
		Amount:         big.NewInt(req.GetAmount()),
	}
}

func toTransferResp(result *processor.TransferResult) *v1.TransferResp {
	return &v1.TransferResp{
		Address: result.TransactionHex,
	}
}

func (tr *TransferHandler) Transfer(ctx context.Context, req *connect_go.Request[v1.TransferReq]) (*connect_go.Response[v1.TransferResp], error) {
	res, err := tr.processor.Transfer(ctx, toTransferParams(req.Msg))
	if err != nil {
		return nil, err
	}
	return connect_go.NewResponse(toTransferResp(res)), nil
}
