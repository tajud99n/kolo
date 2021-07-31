package vault

import (
	"errors"

	"github.com/hashicorp/vault/api"
)

type VaultOptions struct {
	Address string
	Token   string
	Path    string
}

type vault struct {
	logical *api.Logical
	path    string
}


type vaultProviderInterface interface {
	SaveSecrets(map[string]interface{}) error
	FetchSecrets() (map[string]interface{}, error)
}

func VaultConnection(options *VaultOptions) (*vault, error) {
	config := &api.Config{
		Address: options.Address,
	}

	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	client.SetToken(options.Token)

	return &vault{logical: client.Logical(), path: options.Path}, nil
}

// Persist secrets to a path
func (v *vault) SaveSecrets(data map[string]interface{}) error {
	_, err := v.logical.Write(v.path, data)
	if err != nil {
		return err
	}
	return nil
}

// Fetch all the secrets from a path
func (v *vault) FetchSecrets() (map[string]interface{}, error) {
	secrets, err := v.logical.Read(v.path)
	if err != nil {
		return nil, err
	}
	data, ok := secrets.Data["data"].(map[string]interface{})
	if !ok {
		return nil, errors.New("invalid data")
	}
	return data, nil
}
