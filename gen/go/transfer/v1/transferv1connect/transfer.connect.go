// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: transfer/v1/transfer.proto

package transferv1connect

import (
	context "context"
	errors "errors"
	v1 "github.com/avasapollo/eth-erc20/gen/go/transfer/v1"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// TransferServiceName is the fully-qualified name of the TransferService service.
	TransferServiceName = "transfer.v1.TransferService"
)

// TransferServiceClient is a client for the transfer.v1.TransferService service.
type TransferServiceClient interface {
	Transfer(context.Context, *connect_go.Request[v1.TransferReq]) (*connect_go.Response[v1.TransferResp], error)
	Approve(context.Context, *connect_go.Request[v1.ApproveReq]) (*connect_go.Response[v1.ApproveResp], error)
	Allowance(context.Context, *connect_go.Request[v1.AllowanceReq]) (*connect_go.Response[v1.AllowanceResp], error)
	TransferFrom(context.Context, *connect_go.Request[v1.TransferFromReq]) (*connect_go.Response[v1.TransferFromResp], error)
}

// NewTransferServiceClient constructs a client for the transfer.v1.TransferService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTransferServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) TransferServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &transferServiceClient{
		transfer: connect_go.NewClient[v1.TransferReq, v1.TransferResp](
			httpClient,
			baseURL+"/transfer.v1.TransferService/Transfer",
			opts...,
		),
		approve: connect_go.NewClient[v1.ApproveReq, v1.ApproveResp](
			httpClient,
			baseURL+"/transfer.v1.TransferService/Approve",
			opts...,
		),
		allowance: connect_go.NewClient[v1.AllowanceReq, v1.AllowanceResp](
			httpClient,
			baseURL+"/transfer.v1.TransferService/Allowance",
			opts...,
		),
		transferFrom: connect_go.NewClient[v1.TransferFromReq, v1.TransferFromResp](
			httpClient,
			baseURL+"/transfer.v1.TransferService/TransferFrom",
			opts...,
		),
	}
}

// transferServiceClient implements TransferServiceClient.
type transferServiceClient struct {
	transfer     *connect_go.Client[v1.TransferReq, v1.TransferResp]
	approve      *connect_go.Client[v1.ApproveReq, v1.ApproveResp]
	allowance    *connect_go.Client[v1.AllowanceReq, v1.AllowanceResp]
	transferFrom *connect_go.Client[v1.TransferFromReq, v1.TransferFromResp]
}

// Transfer calls transfer.v1.TransferService.Transfer.
func (c *transferServiceClient) Transfer(ctx context.Context, req *connect_go.Request[v1.TransferReq]) (*connect_go.Response[v1.TransferResp], error) {
	return c.transfer.CallUnary(ctx, req)
}

// Approve calls transfer.v1.TransferService.Approve.
func (c *transferServiceClient) Approve(ctx context.Context, req *connect_go.Request[v1.ApproveReq]) (*connect_go.Response[v1.ApproveResp], error) {
	return c.approve.CallUnary(ctx, req)
}

// Allowance calls transfer.v1.TransferService.Allowance.
func (c *transferServiceClient) Allowance(ctx context.Context, req *connect_go.Request[v1.AllowanceReq]) (*connect_go.Response[v1.AllowanceResp], error) {
	return c.allowance.CallUnary(ctx, req)
}

// TransferFrom calls transfer.v1.TransferService.TransferFrom.
func (c *transferServiceClient) TransferFrom(ctx context.Context, req *connect_go.Request[v1.TransferFromReq]) (*connect_go.Response[v1.TransferFromResp], error) {
	return c.transferFrom.CallUnary(ctx, req)
}

// TransferServiceHandler is an implementation of the transfer.v1.TransferService service.
type TransferServiceHandler interface {
	Transfer(context.Context, *connect_go.Request[v1.TransferReq]) (*connect_go.Response[v1.TransferResp], error)
	Approve(context.Context, *connect_go.Request[v1.ApproveReq]) (*connect_go.Response[v1.ApproveResp], error)
	Allowance(context.Context, *connect_go.Request[v1.AllowanceReq]) (*connect_go.Response[v1.AllowanceResp], error)
	TransferFrom(context.Context, *connect_go.Request[v1.TransferFromReq]) (*connect_go.Response[v1.TransferFromResp], error)
}

// NewTransferServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTransferServiceHandler(svc TransferServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/transfer.v1.TransferService/Transfer", connect_go.NewUnaryHandler(
		"/transfer.v1.TransferService/Transfer",
		svc.Transfer,
		opts...,
	))
	mux.Handle("/transfer.v1.TransferService/Approve", connect_go.NewUnaryHandler(
		"/transfer.v1.TransferService/Approve",
		svc.Approve,
		opts...,
	))
	mux.Handle("/transfer.v1.TransferService/Allowance", connect_go.NewUnaryHandler(
		"/transfer.v1.TransferService/Allowance",
		svc.Allowance,
		opts...,
	))
	mux.Handle("/transfer.v1.TransferService/TransferFrom", connect_go.NewUnaryHandler(
		"/transfer.v1.TransferService/TransferFrom",
		svc.TransferFrom,
		opts...,
	))
	return "/transfer.v1.TransferService/", mux
}

// UnimplementedTransferServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTransferServiceHandler struct{}

func (UnimplementedTransferServiceHandler) Transfer(context.Context, *connect_go.Request[v1.TransferReq]) (*connect_go.Response[v1.TransferResp], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("transfer.v1.TransferService.Transfer is not implemented"))
}

func (UnimplementedTransferServiceHandler) Approve(context.Context, *connect_go.Request[v1.ApproveReq]) (*connect_go.Response[v1.ApproveResp], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("transfer.v1.TransferService.Approve is not implemented"))
}

func (UnimplementedTransferServiceHandler) Allowance(context.Context, *connect_go.Request[v1.AllowanceReq]) (*connect_go.Response[v1.AllowanceResp], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("transfer.v1.TransferService.Allowance is not implemented"))
}

func (UnimplementedTransferServiceHandler) TransferFrom(context.Context, *connect_go.Request[v1.TransferFromReq]) (*connect_go.Response[v1.TransferFromResp], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("transfer.v1.TransferService.TransferFrom is not implemented"))
}
