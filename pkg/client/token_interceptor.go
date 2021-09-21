/*
Copyright 2021 CodeNotary, Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package client

import (
	"context"
	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
)

// TokenInterceptor inject token from tokenservice if present and if provided context contain no one
func (c *immuClient) TokenInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	var ctxToken string
	if md, ok := metadata.FromOutgoingContext(ctx); ok && len(md.Get("authorization")) > 0 {
		ctxToken = md.Get("authorization")[0]
	}
	if token, err := c.Tkns.GetToken(); ctxToken == "" && err == nil {
		md := metadata.Pairs("authorization", token)
		ctx = metadata.NewOutgoingContext(ctx, md)
	}
	ris := invoker(ctx, method, req, reply, cc, opts...)

	return ris
}
