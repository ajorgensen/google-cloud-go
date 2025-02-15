// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package retail

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	retailpb "cloud.google.com/go/retail/apiv2/retailpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httpbodypb "google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

var newUserEventClientHook clientHook

// UserEventCallOptions contains the retry settings for each method of UserEventClient.
type UserEventCallOptions struct {
	WriteUserEvent   []gax.CallOption
	CollectUserEvent []gax.CallOption
	PurgeUserEvents  []gax.CallOption
	ImportUserEvents []gax.CallOption
	RejoinUserEvents []gax.CallOption
	GetOperation     []gax.CallOption
	ListOperations   []gax.CallOption
}

func defaultUserEventGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("retail.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("retail.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://retail.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultUserEventCallOptions() *UserEventCallOptions {
	return &UserEventCallOptions{
		WriteUserEvent: []gax.CallOption{
			gax.WithTimeout(5000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        5000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CollectUserEvent: []gax.CallOption{
			gax.WithTimeout(5000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        5000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		PurgeUserEvents: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ImportUserEvents: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        300000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		RejoinUserEvents: []gax.CallOption{
			gax.WithTimeout(5000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        5000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetOperation: []gax.CallOption{},
		ListOperations: []gax.CallOption{
			gax.WithTimeout(300000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        300000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalUserEventClient is an interface that defines the methods available from Retail API.
type internalUserEventClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	WriteUserEvent(context.Context, *retailpb.WriteUserEventRequest, ...gax.CallOption) (*retailpb.UserEvent, error)
	CollectUserEvent(context.Context, *retailpb.CollectUserEventRequest, ...gax.CallOption) (*httpbodypb.HttpBody, error)
	PurgeUserEvents(context.Context, *retailpb.PurgeUserEventsRequest, ...gax.CallOption) (*PurgeUserEventsOperation, error)
	PurgeUserEventsOperation(name string) *PurgeUserEventsOperation
	ImportUserEvents(context.Context, *retailpb.ImportUserEventsRequest, ...gax.CallOption) (*ImportUserEventsOperation, error)
	ImportUserEventsOperation(name string) *ImportUserEventsOperation
	RejoinUserEvents(context.Context, *retailpb.RejoinUserEventsRequest, ...gax.CallOption) (*RejoinUserEventsOperation, error)
	RejoinUserEventsOperation(name string) *RejoinUserEventsOperation
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
	ListOperations(context.Context, *longrunningpb.ListOperationsRequest, ...gax.CallOption) *OperationIterator
}

// UserEventClient is a client for interacting with Retail API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service for ingesting end user actions on the customer website.
type UserEventClient struct {
	// The internal transport-dependent client.
	internalClient internalUserEventClient

	// The call options for this service.
	CallOptions *UserEventCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *UserEventClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *UserEventClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *UserEventClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// WriteUserEvent writes a single user event.
func (c *UserEventClient) WriteUserEvent(ctx context.Context, req *retailpb.WriteUserEventRequest, opts ...gax.CallOption) (*retailpb.UserEvent, error) {
	return c.internalClient.WriteUserEvent(ctx, req, opts...)
}

// CollectUserEvent writes a single user event from the browser. This uses a GET request to
// due to browser restriction of POST-ing to a 3rd party domain.
//
// This method is used only by the Retail API JavaScript pixel and Google Tag
// Manager. Users should not call this method directly.
func (c *UserEventClient) CollectUserEvent(ctx context.Context, req *retailpb.CollectUserEventRequest, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	return c.internalClient.CollectUserEvent(ctx, req, opts...)
}

// PurgeUserEvents deletes permanently all user events specified by the filter provided.
// Depending on the number of events specified by the filter, this operation
// could take hours or days to complete. To test a filter, use the list
// command first.
func (c *UserEventClient) PurgeUserEvents(ctx context.Context, req *retailpb.PurgeUserEventsRequest, opts ...gax.CallOption) (*PurgeUserEventsOperation, error) {
	return c.internalClient.PurgeUserEvents(ctx, req, opts...)
}

// PurgeUserEventsOperation returns a new PurgeUserEventsOperation from a given name.
// The name must be that of a previously created PurgeUserEventsOperation, possibly from a different process.
func (c *UserEventClient) PurgeUserEventsOperation(name string) *PurgeUserEventsOperation {
	return c.internalClient.PurgeUserEventsOperation(name)
}

// ImportUserEvents bulk import of User events. Request processing might be
// synchronous. Events that already exist are skipped.
// Use this method for backfilling historical user events.
//
// Operation.response is of type ImportResponse. Note that it is
// possible for a subset of the items to be successfully inserted.
// Operation.metadata is of type ImportMetadata.
func (c *UserEventClient) ImportUserEvents(ctx context.Context, req *retailpb.ImportUserEventsRequest, opts ...gax.CallOption) (*ImportUserEventsOperation, error) {
	return c.internalClient.ImportUserEvents(ctx, req, opts...)
}

// ImportUserEventsOperation returns a new ImportUserEventsOperation from a given name.
// The name must be that of a previously created ImportUserEventsOperation, possibly from a different process.
func (c *UserEventClient) ImportUserEventsOperation(name string) *ImportUserEventsOperation {
	return c.internalClient.ImportUserEventsOperation(name)
}

// RejoinUserEvents starts a user-event rejoin operation with latest product catalog. Events
// are not annotated with detailed product information for products that are
// missing from the catalog when the user event is ingested. These
// events are stored as unjoined events with limited usage on training and
// serving. You can use this method to start a join operation on specified
// events with the latest version of product catalog. You can also use this
// method to correct events joined with the wrong product catalog. A rejoin
// operation can take hours or days to complete.
func (c *UserEventClient) RejoinUserEvents(ctx context.Context, req *retailpb.RejoinUserEventsRequest, opts ...gax.CallOption) (*RejoinUserEventsOperation, error) {
	return c.internalClient.RejoinUserEvents(ctx, req, opts...)
}

// RejoinUserEventsOperation returns a new RejoinUserEventsOperation from a given name.
// The name must be that of a previously created RejoinUserEventsOperation, possibly from a different process.
func (c *UserEventClient) RejoinUserEventsOperation(name string) *RejoinUserEventsOperation {
	return c.internalClient.RejoinUserEventsOperation(name)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *UserEventClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.GetOperation(ctx, req, opts...)
}

// ListOperations is a utility method from google.longrunning.Operations.
func (c *UserEventClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	return c.internalClient.ListOperations(ctx, req, opts...)
}

// userEventGRPCClient is a client for interacting with Retail API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type userEventGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing UserEventClient
	CallOptions **UserEventCallOptions

	// The gRPC API client.
	userEventClient retailpb.UserEventServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	operationsClient longrunningpb.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewUserEventClient creates a new user event service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service for ingesting end user actions on the customer website.
func NewUserEventClient(ctx context.Context, opts ...option.ClientOption) (*UserEventClient, error) {
	clientOpts := defaultUserEventGRPCClientOptions()
	if newUserEventClientHook != nil {
		hookOpts, err := newUserEventClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := UserEventClient{CallOptions: defaultUserEventCallOptions()}

	c := &userEventGRPCClient{
		connPool:         connPool,
		userEventClient:  retailpb.NewUserEventServiceClient(connPool),
		CallOptions:      &client.CallOptions,
		operationsClient: longrunningpb.NewOperationsClient(connPool),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *userEventGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *userEventGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *userEventGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *userEventGRPCClient) WriteUserEvent(ctx context.Context, req *retailpb.WriteUserEventRequest, opts ...gax.CallOption) (*retailpb.UserEvent, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).WriteUserEvent[0:len((*c.CallOptions).WriteUserEvent):len((*c.CallOptions).WriteUserEvent)], opts...)
	var resp *retailpb.UserEvent
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.userEventClient.WriteUserEvent(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *userEventGRPCClient) CollectUserEvent(ctx context.Context, req *retailpb.CollectUserEventRequest, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CollectUserEvent[0:len((*c.CallOptions).CollectUserEvent):len((*c.CallOptions).CollectUserEvent)], opts...)
	var resp *httpbodypb.HttpBody
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.userEventClient.CollectUserEvent(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *userEventGRPCClient) PurgeUserEvents(ctx context.Context, req *retailpb.PurgeUserEventsRequest, opts ...gax.CallOption) (*PurgeUserEventsOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).PurgeUserEvents[0:len((*c.CallOptions).PurgeUserEvents):len((*c.CallOptions).PurgeUserEvents)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.userEventClient.PurgeUserEvents(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &PurgeUserEventsOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *userEventGRPCClient) ImportUserEvents(ctx context.Context, req *retailpb.ImportUserEventsRequest, opts ...gax.CallOption) (*ImportUserEventsOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ImportUserEvents[0:len((*c.CallOptions).ImportUserEvents):len((*c.CallOptions).ImportUserEvents)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.userEventClient.ImportUserEvents(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &ImportUserEventsOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *userEventGRPCClient) RejoinUserEvents(ctx context.Context, req *retailpb.RejoinUserEventsRequest, opts ...gax.CallOption) (*RejoinUserEventsOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).RejoinUserEvents[0:len((*c.CallOptions).RejoinUserEvents):len((*c.CallOptions).RejoinUserEvents)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.userEventClient.RejoinUserEvents(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &RejoinUserEventsOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *userEventGRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.operationsClient.GetOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *userEventGRPCClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListOperations[0:len((*c.CallOptions).ListOperations):len((*c.CallOptions).ListOperations)], opts...)
	it := &OperationIterator{}
	req = proto.Clone(req).(*longrunningpb.ListOperationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*longrunningpb.Operation, string, error) {
		resp := &longrunningpb.ListOperationsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.operationsClient.ListOperations(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetOperations(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// ImportUserEventsOperation returns a new ImportUserEventsOperation from a given name.
// The name must be that of a previously created ImportUserEventsOperation, possibly from a different process.
func (c *userEventGRPCClient) ImportUserEventsOperation(name string) *ImportUserEventsOperation {
	return &ImportUserEventsOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// PurgeUserEventsOperation returns a new PurgeUserEventsOperation from a given name.
// The name must be that of a previously created PurgeUserEventsOperation, possibly from a different process.
func (c *userEventGRPCClient) PurgeUserEventsOperation(name string) *PurgeUserEventsOperation {
	return &PurgeUserEventsOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// RejoinUserEventsOperation returns a new RejoinUserEventsOperation from a given name.
// The name must be that of a previously created RejoinUserEventsOperation, possibly from a different process.
func (c *userEventGRPCClient) RejoinUserEventsOperation(name string) *RejoinUserEventsOperation {
	return &RejoinUserEventsOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}
