/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package relay

import (
	"context"
	"time"

	"github.com/hyperledger-labs/fabric-smart-client/platform/view/services/flogging"
	"github.com/hyperledger-labs/weaver-dlt-interoperability/common/protos-go/common"
	"github.com/hyperledger-labs/weaver-dlt-interoperability/common/protos-go/relay"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

var logger = flogging.MustGetLogger("fabric-sdk.weaver.relay")

type TimeFunc func() time.Time

type SigningIdentity interface {
	Serialize() ([]byte, error)

	Sign(msg []byte) ([]byte, error)
}

//go:generate counterfeiter -o mock/data_transfer_client.go -fake-name DataTransferClient . DataTransferClient

// DataTransferClient defines an interface that creates a client to communicate with the view service in a peer
type DataTransferClient interface {
	// createDataTransferClient creates a grpc connection and client to the relay server
	CreateDataTransferClient() (*grpc.ClientConn, relay.DataTransferClient, error)
}

func createDataTransferClient(endpoint string) (*grpc.ClientConn, relay.DataTransferClient, error) {
	logger.Infof("opening connection to [%s]", endpoint)
	conn, err := grpc.DialContext(context.Background(), endpoint, grpc.WithInsecure())
	if err != nil {
		logger.Errorf("failed creating connection to [%s]: [%s]", endpoint, err)
		return conn, nil, errors.Wrapf(err, "failed creating connection to [%s]", endpoint)
	}
	logger.Infof("opening connection to [%s], done.", endpoint)

	return conn, relay.NewDataTransferClient(conn), nil
}

type Client struct {}

func (s *Client) RequestState(endpoint string) (*common.Ack, error) {
	conn, client, err := createDataTransferClient(endpoint)
	logger.Infof("get view service client...done")
	if conn != nil {
		logger.Infof("get view service client...got a connection")
		defer conn.Close()
	}
	if err != nil {
		logger.Errorf("failed creating view client [%s]", err)
		return nil, errors.Wrap(err, "failed creating view client")
	}
	ctx := context.Background()
	query := &common.Query{
		Policy:             nil,
		Address:            "",
		RequestingRelay:    "",
		RequestingNetwork:  "",
		Certificate:        "",
		RequestorSignature: "",
		Nonce:              "nonce",
		RequestId:          "",
		RequestingOrg:      "",
	}

	logger.Infof("process request state query [%s]", query.String())
	ack, err := client.RequestState(ctx, query)
	if err != nil {
		logger.Errorf("failed view client process command [%s]", err)
		return nil, errors.Wrap(err, "failed view client process command")
	}
	return ack, nil
}
