// The Package client provides clients that can be used to query servers.
package clients

import (
	c "github.com/alanmvarela/golang-design-patterns/pkg/bridge/credentials_manager"
	log "github.com/alanmvarela/golang-design-patterns/pkg/bridge/logger"
)

// The Client represents a client interface that can be used to query servers.
type Client interface {
	// Query sends a query to the server and returns the response.
	Query(query string) (string, error)
	// Set the logger to use with the client.
	SetLogger(logger log.Logger)
	// SetCredentialsManager sets the credentials manager to use with the client.
	SetCredentialsManager(credentialsManager c.CredentialsManager)
}
