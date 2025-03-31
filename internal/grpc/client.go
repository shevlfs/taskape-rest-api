package grpc

import (
	"context"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"

	pb "taskape-rest-api/proto"
)

// Client represents a gRPC client connection with the backend
type Client struct {
	BackendClient pb.BackendRequestsClient
	conn          *grpc.ClientConn
}

// NewClient creates a new gRPC client with resilient connection settings
func NewClient(backendHost string) (*Client, error) {
	log.Printf("Connecting to backend at %s", backendHost)

	// Create a connection with retry and backoff policies
	conn, err := grpc.Dial(
		backendHost,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDisableRetry(),                              // We'll handle retries ourselves for better control
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)), // Wait for server to be ready
	)
	if err != nil {
		return nil, err
	}

	// Wait up to 5 seconds for initial connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Try to establish initial connection
	go func() {
		for {
			state := conn.GetState()
			if state == connectivity.Ready {
				cancel()
				break
			}
			if !conn.WaitForStateChange(ctx, state) {
				break
			}
		}
	}()

	<-ctx.Done()
	log.Printf("Initial connection state: %v", conn.GetState())

	return &Client{
		BackendClient: pb.NewBackendRequestsClient(conn),
		conn:          conn,
	}, nil
}

// CheckConnection checks if the connection is healthy
func (c *Client) CheckConnection() bool {
	if c.conn == nil {
		return false
	}

	state := c.conn.GetState()
	return state == connectivity.Ready || state == connectivity.Idle
}

// Close closes the gRPC connection
func (c *Client) Close() error {
	if c.conn == nil {
		return nil
	}
	return c.conn.Close()
}

// Reconnect forces a reconnection attempt
func (c *Client) Reconnect(backendHost string) error {
	if c.conn != nil {
		c.conn.Close()
	}

	conn, err := grpc.Dial(
		backendHost,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDisableRetry(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	)
	if err != nil {
		return err
	}

	c.conn = conn
	c.BackendClient = pb.NewBackendRequestsClient(conn)
	return nil
}

// ConnectionManager manages the gRPC client connection with automatic reconnection
type ConnectionManager struct {
	Client      *Client
	backendHost string
	mu          sync.Mutex
	closed      bool
	retryDelay  time.Duration
	maxRetries  int
}

// NewConnectionManager creates a new connection manager
func NewConnectionManager(backendHost string) *ConnectionManager {
	cm := &ConnectionManager{
		backendHost: backendHost,
		retryDelay:  500 * time.Millisecond,
		maxRetries:  10,
	}
	cm.connect()

	// Start health check goroutine
	go cm.healthCheckLoop()

	return cm
}

// connect attempts to connect to the backend
func (cm *ConnectionManager) connect() {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	if cm.closed {
		return
	}

	// Close existing connection if any
	if cm.Client != nil {
		cm.Client.Close()
	}

	var err error
	retries := 0
	backoff := cm.retryDelay

	for retries < cm.maxRetries {
		log.Printf("Connecting to backend (attempt %d/%d)...", retries+1, cm.maxRetries)
		cm.Client, err = NewClient(cm.backendHost)
		if err == nil && cm.Client.CheckConnection() {
			log.Printf("Successfully connected to backend")
			return
		}

		if err != nil {
			log.Printf("Failed to connect: %v. Retrying in %v...", err, backoff)
		} else {
			log.Printf("Connection not ready. Retrying in %v...", backoff)
		}

		time.Sleep(backoff)
		backoff *= 2 // Exponential backoff
		retries++
	}

	log.Printf("Failed to establish connection after %d attempts", cm.maxRetries)
}

// GetClient returns the current client, connecting if necessary
func (cm *ConnectionManager) GetClient() *Client {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	if cm.Client == nil || !cm.Client.CheckConnection() {
		// Release lock during connection attempt
		cm.mu.Unlock()
		cm.connect()
		cm.mu.Lock()
	}

	return cm.Client
}

// healthCheckLoop periodically checks the connection health
func (cm *ConnectionManager) healthCheckLoop() {
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if cm.closed {
			return
		}

		cm.mu.Lock()
		healthy := cm.Client != nil && cm.Client.CheckConnection()
		cm.mu.Unlock()

		if !healthy {
			log.Println("Connection lost or unhealthy, attempting to reconnect...")
			cm.connect()
		}
	}
}

// Close closes the connection manager and its client
func (cm *ConnectionManager) Close() error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	cm.closed = true
	if cm.Client != nil {
		return cm.Client.Close()
	}

	return nil
}
