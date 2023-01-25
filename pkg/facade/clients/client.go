// Package client provides client functions to send queries to different APIs
package clients

// Client is the interface for every implemented client.
type Client interface {
	Query() (*string, error)
}
