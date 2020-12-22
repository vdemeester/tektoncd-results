// Copyright 2020 The Tekton Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package test

import (
	"context"
	"fmt"
	"net"
	"testing"

	ttesting "github.com/tektoncd/pipeline/pkg/reconciler/testing"
	pipelinetest "github.com/tektoncd/pipeline/test"
	"github.com/tektoncd/results/pkg/api/server/test"
	v1alpha1server "github.com/tektoncd/results/pkg/api/server/v1alpha1"
	server "github.com/tektoncd/results/pkg/api/server/v1alpha2"
	v1alpha1pb "github.com/tektoncd/results/proto/v1alpha1/results_go_proto"
	pb "github.com/tektoncd/results/proto/v1alpha2/results_go_proto"
	"google.golang.org/grpc"
	"knative.dev/pkg/configmap"
)

const (
	port = ":0"
)

func NewResultsClient(t *testing.T) pb.ResultsClient {
	t.Helper()
	srv, err := server.New(test.NewDB(t))
	if err != nil {
		t.Fatalf("Failed to create fake server: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterResultsServer(s, srv) // local test server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}
	go func() {
		if err := s.Serve(lis); err != nil {
			fmt.Printf("error starting result server: %v\n", err)
		}
	}()
	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	t.Cleanup(func() {
		s.Stop()
		lis.Close()
		conn.Close()
	})
	return pb.NewResultsClient(conn)
}

func NewLegacyResultsClient(t *testing.T) v1alpha1pb.ResultsClient {
	t.Helper()
	srv, err := v1alpha1server.New(test.NewDB(t))
	if err != nil {
		t.Fatalf("Failed to create fake server: %v", err)
	}
	s := grpc.NewServer()
	v1alpha1pb.RegisterResultsServer(s, srv) // local test server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}
	go func() {
		if err := s.Serve(lis); err != nil {
			fmt.Printf("error starting result server: %v\n", err)
		}
	}()
	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	t.Cleanup(func() {
		s.Stop()
		lis.Close()
		conn.Close()
	})
	return v1alpha1pb.NewResultsClient(conn)
}

func GetFakeClients(t *testing.T, d pipelinetest.Data, client v1alpha1pb.ResultsClient) (context.Context, pipelinetest.Clients, *configmap.InformedWatcher) {
	t.Helper()
	ctx, _ := ttesting.SetupFakeContext(t)
	clients, _ := pipelinetest.SeedTestData(t, ctx, d)
	cmw := configmap.NewInformedWatcher(clients.Kube, "")
	return ctx, clients, cmw
}