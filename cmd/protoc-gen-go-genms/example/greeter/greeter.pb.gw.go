// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: greeter.proto

/*
Package greeter is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package greeter

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_WithRest_HelloRest_0(ctx context.Context, marshaler runtime.Marshaler, client WithRestClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Message
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.HelloRest(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_WithRest_HelloRest_0(ctx context.Context, marshaler runtime.Marshaler, server WithRestServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Message
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.HelloRest(ctx, &protoReq)
	return msg, metadata, err

}

func request_WithGraphQL_HelloGraphQL_0(ctx context.Context, marshaler runtime.Marshaler, client WithGraphQLClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Message
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.HelloGraphQL(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_WithGraphQL_HelloGraphQL_0(ctx context.Context, marshaler runtime.Marshaler, server WithGraphQLServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Message
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.HelloGraphQL(ctx, &protoReq)
	return msg, metadata, err

}

func request_WithRestAndGraphQL_HelloRestAndGraphQL_0(ctx context.Context, marshaler runtime.Marshaler, client WithRestAndGraphQLClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Message
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.HelloRestAndGraphQL(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_WithRestAndGraphQL_HelloRestAndGraphQL_0(ctx context.Context, marshaler runtime.Marshaler, server WithRestAndGraphQLServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Message
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.HelloRestAndGraphQL(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterWithRestHandlerServer registers the http handlers for service WithRest to "mux".
// UnaryRPC     :call WithRestServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterWithRestHandlerFromEndpoint instead.
func RegisterWithRestHandlerServer(ctx context.Context, mux *runtime.ServeMux, server WithRestServer) error {

	mux.Handle("POST", pattern_WithRest_HelloRest_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/greeter.WithRest/HelloRest")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_WithRest_HelloRest_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_WithRest_HelloRest_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterWithGraphQLHandlerServer registers the http handlers for service WithGraphQL to "mux".
// UnaryRPC     :call WithGraphQLServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterWithGraphQLHandlerFromEndpoint instead.
func RegisterWithGraphQLHandlerServer(ctx context.Context, mux *runtime.ServeMux, server WithGraphQLServer) error {

	mux.Handle("POST", pattern_WithGraphQL_HelloGraphQL_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/greeter.WithGraphQL/HelloGraphQL")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_WithGraphQL_HelloGraphQL_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_WithGraphQL_HelloGraphQL_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterWithRestAndGraphQLHandlerServer registers the http handlers for service WithRestAndGraphQL to "mux".
// UnaryRPC     :call WithRestAndGraphQLServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterWithRestAndGraphQLHandlerFromEndpoint instead.
func RegisterWithRestAndGraphQLHandlerServer(ctx context.Context, mux *runtime.ServeMux, server WithRestAndGraphQLServer) error {

	mux.Handle("POST", pattern_WithRestAndGraphQL_HelloRestAndGraphQL_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/greeter.WithRestAndGraphQL/HelloRestAndGraphQL")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_WithRestAndGraphQL_HelloRestAndGraphQL_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_WithRestAndGraphQL_HelloRestAndGraphQL_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterWithRestHandlerFromEndpoint is same as RegisterWithRestHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterWithRestHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterWithRestHandler(ctx, mux, conn)
}

// RegisterWithRestHandler registers the http handlers for service WithRest to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterWithRestHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterWithRestHandlerClient(ctx, mux, NewWithRestClient(conn))
}

// RegisterWithRestHandlerClient registers the http handlers for service WithRest
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "WithRestClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "WithRestClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "WithRestClient" to call the correct interceptors.
func RegisterWithRestHandlerClient(ctx context.Context, mux *runtime.ServeMux, client WithRestClient) error {

	mux.Handle("POST", pattern_WithRest_HelloRest_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/greeter.WithRest/HelloRest")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_WithRest_HelloRest_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_WithRest_HelloRest_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_WithRest_HelloRest_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "rest"}, ""))
)

var (
	forward_WithRest_HelloRest_0 = runtime.ForwardResponseMessage
)

// RegisterWithGraphQLHandlerFromEndpoint is same as RegisterWithGraphQLHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterWithGraphQLHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterWithGraphQLHandler(ctx, mux, conn)
}

// RegisterWithGraphQLHandler registers the http handlers for service WithGraphQL to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterWithGraphQLHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterWithGraphQLHandlerClient(ctx, mux, NewWithGraphQLClient(conn))
}

// RegisterWithGraphQLHandlerClient registers the http handlers for service WithGraphQL
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "WithGraphQLClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "WithGraphQLClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "WithGraphQLClient" to call the correct interceptors.
func RegisterWithGraphQLHandlerClient(ctx context.Context, mux *runtime.ServeMux, client WithGraphQLClient) error {

	mux.Handle("POST", pattern_WithGraphQL_HelloGraphQL_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/greeter.WithGraphQL/HelloGraphQL")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_WithGraphQL_HelloGraphQL_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_WithGraphQL_HelloGraphQL_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_WithGraphQL_HelloGraphQL_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "graphql"}, ""))
)

var (
	forward_WithGraphQL_HelloGraphQL_0 = runtime.ForwardResponseMessage
)

// RegisterWithRestAndGraphQLHandlerFromEndpoint is same as RegisterWithRestAndGraphQLHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterWithRestAndGraphQLHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterWithRestAndGraphQLHandler(ctx, mux, conn)
}

// RegisterWithRestAndGraphQLHandler registers the http handlers for service WithRestAndGraphQL to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterWithRestAndGraphQLHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterWithRestAndGraphQLHandlerClient(ctx, mux, NewWithRestAndGraphQLClient(conn))
}

// RegisterWithRestAndGraphQLHandlerClient registers the http handlers for service WithRestAndGraphQL
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "WithRestAndGraphQLClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "WithRestAndGraphQLClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "WithRestAndGraphQLClient" to call the correct interceptors.
func RegisterWithRestAndGraphQLHandlerClient(ctx context.Context, mux *runtime.ServeMux, client WithRestAndGraphQLClient) error {

	mux.Handle("POST", pattern_WithRestAndGraphQL_HelloRestAndGraphQL_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/greeter.WithRestAndGraphQL/HelloRestAndGraphQL")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_WithRestAndGraphQL_HelloRestAndGraphQL_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_WithRestAndGraphQL_HelloRestAndGraphQL_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_WithRestAndGraphQL_HelloRestAndGraphQL_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "rest", "graphql"}, ""))
)

var (
	forward_WithRestAndGraphQL_HelloRestAndGraphQL_0 = runtime.ForwardResponseMessage
)
