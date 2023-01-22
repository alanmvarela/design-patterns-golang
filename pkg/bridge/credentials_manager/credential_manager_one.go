// Package credentialsmanager provides a bridge to the credentials manager.
package credentialsmanager

import "errors"

// CredentialsManagerOne is a implementation of the credentials manager interface.
type CredentialsManagerOne struct {
}

// GetCredentials returns the credentials from the credential manager.
func (c *CredentialsManagerOne) GetCredentials(account string) (*string, error) {
	if account == "" {
		return nil, errors.New("account cannot be empty")
	}
	credentials := "Credentials for account " + account + " from credential manager One"
	if credentials == "" {
		return nil, errors.New("credentials cannot be empty")
	}
	return &credentials, nil
}
