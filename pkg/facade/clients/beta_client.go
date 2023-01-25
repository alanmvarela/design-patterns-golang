// Package client provides client functions to send queries to different APIs
package clients

type BetaClient struct {
}

// Query sends a query to the Beta API
func (c *BetaClient) Query() (*string, error) {
	response := "Beta client response"
	return &response, nil
}
