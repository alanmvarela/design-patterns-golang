// The Package adapter provides the adapter implementations to use the third party clients.
package adapter

// The Client interface defines the methods that all clients should implement.
// The adapters in this package will have to implement this interface.
type DNSClient interface {
	// Query sends a query to the server and returns the response.
	Query(q string) (*string, error)
}
