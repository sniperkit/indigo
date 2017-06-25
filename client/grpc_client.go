//  Copyright (c) 2017 Minoru Osuka
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"github.com/mosuka/indigo/proto"
	"google.golang.org/grpc"
)

type indigoGRPCClient struct {
	Server     string
	Connection *grpc.ClientConn
	Client     proto.IndigoClient
}

func NewIndigoGRPCClient(server string) (*indigoGRPCClient, error) {
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		return &indigoGRPCClient{}, err
	}

	ic := proto.NewIndigoClient(conn)

	return &indigoGRPCClient{
		Server:     server,
		Connection: conn,
		Client:     ic,
	}, nil
}

func (igc *indigoGRPCClient) Close() error {
	err := igc.Connection.Close()
	if err != nil {
		return err
	}

	return nil
}
