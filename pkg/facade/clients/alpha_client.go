// Package client provides client functions to send queries to different APIs
package clients

type AlphaClient struct {
}

// Query sends a query to the Alpha API
func (c *AlphaClient) Query() (*string, error) {
	response := "Alpha client response"
	return &response, nil
}
