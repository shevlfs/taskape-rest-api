package grpc

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "taskape-rest-api/proto"
)

type Client struct {
	BackendClient pb.BackendRequestsClient
	conn          *grpc.ClientConn
}

func NewClient(backendHost string) (*Client, error) {
	log.Printf("Connecting to backend at %s", backendHost)

	conn, err := grpc.Dial(backendHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Client{
		BackendClient: pb.NewBackendRequestsClient(conn),
		conn:          conn,
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}
