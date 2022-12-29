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
	Approve(ctx context.Context, req *processor.ApproveParams) (*processor.ApproveResult, error)
	Allowance(ctx context.Context, req *processor.AllowanceParams) (*processor.AllowanceResult, error)
	TransferFrom(ctx context.Context, req *processor.TransferFromParams) (*processor.TransferFromResult, error)
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

func toApproveParams(req *v1.ApproveReq) *processor.ApproveParams {
	return &processor.ApproveParams{
		OwnerAddr:     req.GetOwnerAddress(),
		OwnerPassword: req.GetOwnerPassword(),
		DelegateAddr:  req.GetDelegateAddress(),
		Amount:        big.NewInt(req.GetAmount()),
	}
}

func toApproveRsp(result *processor.ApproveResult) *v1.ApproveResp {
	return &v1.ApproveResp{Address: result.TransactionHex}
}

func (tr *TransferHandler) Approve(ctx context.Context, req *connect_go.Request[v1.ApproveReq]) (*connect_go.Response[v1.ApproveResp], error) {
	res, err := tr.processor.Approve(ctx, toApproveParams(req.Msg))
	if err != nil {
		return nil, err
	}
	return connect_go.NewResponse(toApproveRsp(res)), nil
}

func toAllowanceParams(req *v1.AllowanceReq) *processor.AllowanceParams {
	return &processor.AllowanceParams{
		OwnerAddr:    req.GetOwnerAddress(),
		DelegateAddr: req.GetDelegateAddress(),
	}
}

func toAllowanceResp(result *processor.AllowanceResult) *v1.AllowanceResp {
	return &v1.AllowanceResp{
		Amount: result.Amount,
	}
}

func (tr *TransferHandler) Allowance(ctx context.Context, req *connect_go.Request[v1.AllowanceReq]) (*connect_go.Response[v1.AllowanceResp], error) {
	res, err := tr.processor.Allowance(ctx, toAllowanceParams(req.Msg))
	if err != nil {
		return nil, err
	}
	return connect_go.NewResponse(toAllowanceResp(res)), nil
}

func toTransferFromParams(req *v1.TransferFromReq) *processor.TransferFromParams {
	return &processor.TransferFromParams{
		DelegateAddr:     req.GetDelegateAddress(),
		DelegatePassword: req.GetDelegatePassword(),
		OwnerAddr:        req.GetOwnerAddress(),
		BuyerAddr:        req.GetBuyerAddress(),
		Amount:           big.NewInt(req.Amount),
	}
}

func toTransferFromResp(result *processor.TransferFromResult) *v1.TransferFromResp {
	return &v1.TransferFromResp{
		Address: result.TransactionHex,
	}
}

func (tr *TransferHandler) TransferFrom(ctx context.Context, req *connect_go.Request[v1.TransferFromReq]) (*connect_go.Response[v1.TransferFromResp], error) {
	res, err := tr.processor.TransferFrom(ctx, toTransferFromParams(req.Msg))
	if err != nil {
		return nil, err
	}
	return connect_go.NewResponse(toTransferFromResp(res)), nil
}
