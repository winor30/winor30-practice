// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: buf/starwars/starship/v1/starship.proto

package starshipv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/bufbuild/knit/tutorial/starwars-knit-relation-service-go/gen/buf/starwars/starship/v1"
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
	// StarshipServiceName is the fully-qualified name of the StarshipService service.
	StarshipServiceName = "buf.starwars.starship.v1.StarshipService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// StarshipServiceGetStarshipsProcedure is the fully-qualified name of the StarshipService's
	// GetStarships RPC.
	StarshipServiceGetStarshipsProcedure = "/buf.starwars.starship.v1.StarshipService/GetStarships"
)

// StarshipServiceClient is a client for the buf.starwars.starship.v1.StarshipService service.
type StarshipServiceClient interface {
	GetStarships(context.Context, *connect_go.Request[v1.StarshipsRequest]) (*connect_go.Response[v1.StarshipsResponse], error)
}

// NewStarshipServiceClient constructs a client for the buf.starwars.starship.v1.StarshipService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewStarshipServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) StarshipServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &starshipServiceClient{
		getStarships: connect_go.NewClient[v1.StarshipsRequest, v1.StarshipsResponse](
			httpClient,
			baseURL+StarshipServiceGetStarshipsProcedure,
			opts...,
		),
	}
}

// starshipServiceClient implements StarshipServiceClient.
type starshipServiceClient struct {
	getStarships *connect_go.Client[v1.StarshipsRequest, v1.StarshipsResponse]
}

// GetStarships calls buf.starwars.starship.v1.StarshipService.GetStarships.
func (c *starshipServiceClient) GetStarships(ctx context.Context, req *connect_go.Request[v1.StarshipsRequest]) (*connect_go.Response[v1.StarshipsResponse], error) {
	return c.getStarships.CallUnary(ctx, req)
}

// StarshipServiceHandler is an implementation of the buf.starwars.starship.v1.StarshipService
// service.
type StarshipServiceHandler interface {
	GetStarships(context.Context, *connect_go.Request[v1.StarshipsRequest]) (*connect_go.Response[v1.StarshipsResponse], error)
}

// NewStarshipServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewStarshipServiceHandler(svc StarshipServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(StarshipServiceGetStarshipsProcedure, connect_go.NewUnaryHandler(
		StarshipServiceGetStarshipsProcedure,
		svc.GetStarships,
		opts...,
	))
	return "/buf.starwars.starship.v1.StarshipService/", mux
}

// UnimplementedStarshipServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedStarshipServiceHandler struct{}

func (UnimplementedStarshipServiceHandler) GetStarships(context.Context, *connect_go.Request[v1.StarshipsRequest]) (*connect_go.Response[v1.StarshipsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.starwars.starship.v1.StarshipService.GetStarships is not implemented"))
}