// Package credentialsmanager provides a bridge to the credentials manager.
package credentialsmanager

type CredentialsManager interface {
	// GetCredentials returns the credentials for the given account.
	GetCredentials(account string) (*string, error)
}
