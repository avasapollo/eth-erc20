package handlers

import (
	"context"

	v1 "github.com/avasapollo/eth-erc20/gen/go/balance/v1"
	"github.com/avasapollo/eth-erc20/gen/go/balance/v1/balancev1connect"
	"github.com/avasapollo/eth-erc20/processor"
	connect_go "github.com/bufbuild/connect-go"
)

var _ balancev1connect.BalanceServiceHandler = (*BalanceHandler)(nil)

type BalanceService interface {
	GetBalance(ctx context.Context, address string) (*processor.GetBalanceResult, error)
}

type BalanceHandler struct {
	svc BalanceService
}

func NewBalanceHandler(svc BalanceService) *BalanceHandler {
	return &BalanceHandler{
		svc: svc,
	}
}

func toGetBalanceResp(res *processor.GetBalanceResult) *v1.GetBalanceResp {
	return &v1.GetBalanceResp{
		Amount: res.Amount,
	}
}

func (b *BalanceHandler) GetBalance(ctx context.Context, req *connect_go.Request[v1.GetBalanceReq]) (*connect_go.Response[v1.GetBalanceResp], error) {
	res, err := b.svc.GetBalance(ctx, req.Msg.GetAddress())
	if err != nil {
		return nil, err
	}
	return connect_go.NewResponse(toGetBalanceResp(res)), err
}
