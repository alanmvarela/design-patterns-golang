// Package credentialsmanager provides a bridge to the credentials manager.
package credentialsmanager

import "errors"

// CredentialsManagerTwo is a implementation of the credentials manager interface.
type CredentialsManagerTwo struct {
}

// GetCredentials returns the credentials from the credential manager.
func (c *CredentialsManagerTwo) GetCredentials(account string) (*string, error) {
	if account == "" {
		return nil, errors.New("account cannot be empty")
	}
	credentials := "Credentials for account " + account + " from credential manager Two"
	if credentials == "" {
		return nil, errors.New("credentials cannot be empty")
	}
	return &credentials, nil
}
