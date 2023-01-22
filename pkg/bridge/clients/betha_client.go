// The Package client provides clients that can be used to query servers.
package clients

import (
	"errors"

	c "github.com/alanmvarela/golang-design-patterns/pkg/bridge/credentials_manager"
	log "github.com/alanmvarela/golang-design-patterns/pkg/bridge/logger"
)

type BethaClient struct {
	// The logger to use with the client.
	logger log.Logger
	// The credentials manager to use with the client.
	credentialsManager c.CredentialsManager
}

// NewBethaClient returns a new instance of the BethaClient.
func NewBethaClient() *BethaClient {
	return &BethaClient{}
}

// Query sends a query to the server and returns the response.
func (c *BethaClient) Query(query string) (*string, error) {
	// Using log and credentials manager from the bridge.
	c.logger.Log("Querying Betha API with query " + query)
	credentials, error := c.credentialsManager.GetCredentials("betha")
	if error != nil {
		return nil, error
	}
	if query == "" {
		return nil, errors.New("query cannot be empty")
	}
	response := "Response from Betha API for query " + query + " with credentials: " + *credentials
	return &response, nil
}

// SetLogger sets the logger to use with the client.
func (c *BethaClient) SetLogger(logger log.Logger) {
	c.logger = logger
}

// SetCredentialsManager sets the credentials manager to use with the client.
func (c *BethaClient) SetCredentialsManager(credentialsManager c.CredentialsManager) {
	c.credentialsManager = credentialsManager
}
