// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: envd/process/v1/process.proto

package processv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/e2b-dev/infra/packages/envd/internal/services/spec/envd/process/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// ProcessServiceName is the fully-qualified name of the ProcessService service.
	ProcessServiceName = "envd.process.v1.ProcessService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ProcessServiceListProcessesProcedure is the fully-qualified name of the ProcessService's
	// ListProcesses RPC.
	ProcessServiceListProcessesProcedure = "/envd.process.v1.ProcessService/ListProcesses"
	// ProcessServiceReconnectProcessProcedure is the fully-qualified name of the ProcessService's
	// ReconnectProcess RPC.
	ProcessServiceReconnectProcessProcedure = "/envd.process.v1.ProcessService/ReconnectProcess"
	// ProcessServiceStartProcessProcedure is the fully-qualified name of the ProcessService's
	// StartProcess RPC.
	ProcessServiceStartProcessProcedure = "/envd.process.v1.ProcessService/StartProcess"
	// ProcessServiceUpdateProcessProcedure is the fully-qualified name of the ProcessService's
	// UpdateProcess RPC.
	ProcessServiceUpdateProcessProcedure = "/envd.process.v1.ProcessService/UpdateProcess"
	// ProcessServiceSendProcessInputProcedure is the fully-qualified name of the ProcessService's
	// SendProcessInput RPC.
	ProcessServiceSendProcessInputProcedure = "/envd.process.v1.ProcessService/SendProcessInput"
	// ProcessServiceSendProcessSignalProcedure is the fully-qualified name of the ProcessService's
	// SendProcessSignal RPC.
	ProcessServiceSendProcessSignalProcedure = "/envd.process.v1.ProcessService/SendProcessSignal"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	processServiceServiceDescriptor                 = v1.File_envd_process_v1_process_proto.Services().ByName("ProcessService")
	processServiceListProcessesMethodDescriptor     = processServiceServiceDescriptor.Methods().ByName("ListProcesses")
	processServiceReconnectProcessMethodDescriptor  = processServiceServiceDescriptor.Methods().ByName("ReconnectProcess")
	processServiceStartProcessMethodDescriptor      = processServiceServiceDescriptor.Methods().ByName("StartProcess")
	processServiceUpdateProcessMethodDescriptor     = processServiceServiceDescriptor.Methods().ByName("UpdateProcess")
	processServiceSendProcessInputMethodDescriptor  = processServiceServiceDescriptor.Methods().ByName("SendProcessInput")
	processServiceSendProcessSignalMethodDescriptor = processServiceServiceDescriptor.Methods().ByName("SendProcessSignal")
)

// ProcessServiceClient is a client for the envd.process.v1.ProcessService service.
type ProcessServiceClient interface {
	ListProcesses(context.Context, *connect.Request[v1.ListProcessesRequest]) (*connect.Response[v1.ListProcessesResponse], error)
	ReconnectProcess(context.Context, *connect.Request[v1.ReconnectProcessRequest]) (*connect.ServerStreamForClient[v1.ReconnectProcessResponse], error)
	StartProcess(context.Context, *connect.Request[v1.StartProcessRequest]) (*connect.ServerStreamForClient[v1.StartProcessResponse], error)
	UpdateProcess(context.Context, *connect.Request[v1.UpdateProcessRequest]) (*connect.Response[v1.UpdateProcessResponse], error)
	SendProcessInput(context.Context, *connect.Request[v1.SendProcessInputRequest]) (*connect.Response[v1.SendProcessInputResponse], error)
	SendProcessSignal(context.Context, *connect.Request[v1.SendProcessSignalRequest]) (*connect.Response[v1.SendProcessSignalResponse], error)
}

// NewProcessServiceClient constructs a client for the envd.process.v1.ProcessService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewProcessServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ProcessServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &processServiceClient{
		listProcesses: connect.NewClient[v1.ListProcessesRequest, v1.ListProcessesResponse](
			httpClient,
			baseURL+ProcessServiceListProcessesProcedure,
			connect.WithSchema(processServiceListProcessesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		reconnectProcess: connect.NewClient[v1.ReconnectProcessRequest, v1.ReconnectProcessResponse](
			httpClient,
			baseURL+ProcessServiceReconnectProcessProcedure,
			connect.WithSchema(processServiceReconnectProcessMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		startProcess: connect.NewClient[v1.StartProcessRequest, v1.StartProcessResponse](
			httpClient,
			baseURL+ProcessServiceStartProcessProcedure,
			connect.WithSchema(processServiceStartProcessMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateProcess: connect.NewClient[v1.UpdateProcessRequest, v1.UpdateProcessResponse](
			httpClient,
			baseURL+ProcessServiceUpdateProcessProcedure,
			connect.WithSchema(processServiceUpdateProcessMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		sendProcessInput: connect.NewClient[v1.SendProcessInputRequest, v1.SendProcessInputResponse](
			httpClient,
			baseURL+ProcessServiceSendProcessInputProcedure,
			connect.WithSchema(processServiceSendProcessInputMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		sendProcessSignal: connect.NewClient[v1.SendProcessSignalRequest, v1.SendProcessSignalResponse](
			httpClient,
			baseURL+ProcessServiceSendProcessSignalProcedure,
			connect.WithSchema(processServiceSendProcessSignalMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// processServiceClient implements ProcessServiceClient.
type processServiceClient struct {
	listProcesses     *connect.Client[v1.ListProcessesRequest, v1.ListProcessesResponse]
	reconnectProcess  *connect.Client[v1.ReconnectProcessRequest, v1.ReconnectProcessResponse]
	startProcess      *connect.Client[v1.StartProcessRequest, v1.StartProcessResponse]
	updateProcess     *connect.Client[v1.UpdateProcessRequest, v1.UpdateProcessResponse]
	sendProcessInput  *connect.Client[v1.SendProcessInputRequest, v1.SendProcessInputResponse]
	sendProcessSignal *connect.Client[v1.SendProcessSignalRequest, v1.SendProcessSignalResponse]
}

// ListProcesses calls envd.process.v1.ProcessService.ListProcesses.
func (c *processServiceClient) ListProcesses(ctx context.Context, req *connect.Request[v1.ListProcessesRequest]) (*connect.Response[v1.ListProcessesResponse], error) {
	return c.listProcesses.CallUnary(ctx, req)
}

// ReconnectProcess calls envd.process.v1.ProcessService.ReconnectProcess.
func (c *processServiceClient) ReconnectProcess(ctx context.Context, req *connect.Request[v1.ReconnectProcessRequest]) (*connect.ServerStreamForClient[v1.ReconnectProcessResponse], error) {
	return c.reconnectProcess.CallServerStream(ctx, req)
}

// StartProcess calls envd.process.v1.ProcessService.StartProcess.
func (c *processServiceClient) StartProcess(ctx context.Context, req *connect.Request[v1.StartProcessRequest]) (*connect.ServerStreamForClient[v1.StartProcessResponse], error) {
	return c.startProcess.CallServerStream(ctx, req)
}

// UpdateProcess calls envd.process.v1.ProcessService.UpdateProcess.
func (c *processServiceClient) UpdateProcess(ctx context.Context, req *connect.Request[v1.UpdateProcessRequest]) (*connect.Response[v1.UpdateProcessResponse], error) {
	return c.updateProcess.CallUnary(ctx, req)
}

// SendProcessInput calls envd.process.v1.ProcessService.SendProcessInput.
func (c *processServiceClient) SendProcessInput(ctx context.Context, req *connect.Request[v1.SendProcessInputRequest]) (*connect.Response[v1.SendProcessInputResponse], error) {
	return c.sendProcessInput.CallUnary(ctx, req)
}

// SendProcessSignal calls envd.process.v1.ProcessService.SendProcessSignal.
func (c *processServiceClient) SendProcessSignal(ctx context.Context, req *connect.Request[v1.SendProcessSignalRequest]) (*connect.Response[v1.SendProcessSignalResponse], error) {
	return c.sendProcessSignal.CallUnary(ctx, req)
}

// ProcessServiceHandler is an implementation of the envd.process.v1.ProcessService service.
type ProcessServiceHandler interface {
	ListProcesses(context.Context, *connect.Request[v1.ListProcessesRequest]) (*connect.Response[v1.ListProcessesResponse], error)
	ReconnectProcess(context.Context, *connect.Request[v1.ReconnectProcessRequest], *connect.ServerStream[v1.ReconnectProcessResponse]) error
	StartProcess(context.Context, *connect.Request[v1.StartProcessRequest], *connect.ServerStream[v1.StartProcessResponse]) error
	UpdateProcess(context.Context, *connect.Request[v1.UpdateProcessRequest]) (*connect.Response[v1.UpdateProcessResponse], error)
	SendProcessInput(context.Context, *connect.Request[v1.SendProcessInputRequest]) (*connect.Response[v1.SendProcessInputResponse], error)
	SendProcessSignal(context.Context, *connect.Request[v1.SendProcessSignalRequest]) (*connect.Response[v1.SendProcessSignalResponse], error)
}

// NewProcessServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewProcessServiceHandler(svc ProcessServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	processServiceListProcessesHandler := connect.NewUnaryHandler(
		ProcessServiceListProcessesProcedure,
		svc.ListProcesses,
		connect.WithSchema(processServiceListProcessesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	processServiceReconnectProcessHandler := connect.NewServerStreamHandler(
		ProcessServiceReconnectProcessProcedure,
		svc.ReconnectProcess,
		connect.WithSchema(processServiceReconnectProcessMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	processServiceStartProcessHandler := connect.NewServerStreamHandler(
		ProcessServiceStartProcessProcedure,
		svc.StartProcess,
		connect.WithSchema(processServiceStartProcessMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	processServiceUpdateProcessHandler := connect.NewUnaryHandler(
		ProcessServiceUpdateProcessProcedure,
		svc.UpdateProcess,
		connect.WithSchema(processServiceUpdateProcessMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	processServiceSendProcessInputHandler := connect.NewUnaryHandler(
		ProcessServiceSendProcessInputProcedure,
		svc.SendProcessInput,
		connect.WithSchema(processServiceSendProcessInputMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	processServiceSendProcessSignalHandler := connect.NewUnaryHandler(
		ProcessServiceSendProcessSignalProcedure,
		svc.SendProcessSignal,
		connect.WithSchema(processServiceSendProcessSignalMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/envd.process.v1.ProcessService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ProcessServiceListProcessesProcedure:
			processServiceListProcessesHandler.ServeHTTP(w, r)
		case ProcessServiceReconnectProcessProcedure:
			processServiceReconnectProcessHandler.ServeHTTP(w, r)
		case ProcessServiceStartProcessProcedure:
			processServiceStartProcessHandler.ServeHTTP(w, r)
		case ProcessServiceUpdateProcessProcedure:
			processServiceUpdateProcessHandler.ServeHTTP(w, r)
		case ProcessServiceSendProcessInputProcedure:
			processServiceSendProcessInputHandler.ServeHTTP(w, r)
		case ProcessServiceSendProcessSignalProcedure:
			processServiceSendProcessSignalHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedProcessServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedProcessServiceHandler struct{}

func (UnimplementedProcessServiceHandler) ListProcesses(context.Context, *connect.Request[v1.ListProcessesRequest]) (*connect.Response[v1.ListProcessesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("envd.process.v1.ProcessService.ListProcesses is not implemented"))
}

func (UnimplementedProcessServiceHandler) ReconnectProcess(context.Context, *connect.Request[v1.ReconnectProcessRequest], *connect.ServerStream[v1.ReconnectProcessResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("envd.process.v1.ProcessService.ReconnectProcess is not implemented"))
}

func (UnimplementedProcessServiceHandler) StartProcess(context.Context, *connect.Request[v1.StartProcessRequest], *connect.ServerStream[v1.StartProcessResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("envd.process.v1.ProcessService.StartProcess is not implemented"))
}

func (UnimplementedProcessServiceHandler) UpdateProcess(context.Context, *connect.Request[v1.UpdateProcessRequest]) (*connect.Response[v1.UpdateProcessResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("envd.process.v1.ProcessService.UpdateProcess is not implemented"))
}

func (UnimplementedProcessServiceHandler) SendProcessInput(context.Context, *connect.Request[v1.SendProcessInputRequest]) (*connect.Response[v1.SendProcessInputResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("envd.process.v1.ProcessService.SendProcessInput is not implemented"))
}

func (UnimplementedProcessServiceHandler) SendProcessSignal(context.Context, *connect.Request[v1.SendProcessSignalRequest]) (*connect.Response[v1.SendProcessSignalResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("envd.process.v1.ProcessService.SendProcessSignal is not implemented"))
}
