package client

import (
	"time"

	"github.com/IcaroSilvaFK/go_lang_infiscal/cmd/helpers"
)

type ClientCredentials struct {
	ServiceTokenKey string
}

type WorkspaceConfig struct {
	WorkspaceId  string
	Environment  string
	WorkspaceKey string
}

type SecretBundle struct {
	SecretName    string
	SecretValue   string
	Version       int
	Workspace     string
	Environment   string
	Type          string
	IsFallback    bool
	LastFetchedAt time.Time
}

type InfiscalClient struct {
	AuthMode    string
	Credentials ClientCredentials
	WorkspaceConfig
}

type InfiscalClientConfig interface {
	GetAllSecrets() map[string]string
	GetSecret(name string, options string) SecretBundle
	CreateSecret(name, value string) SecretBundle
	DeleteSecret(name string)
}

func NewInfiscalClient(
	authMode string,
	credentials ClientCredentials,
	workspaceConfig WorkspaceConfig,
) InfiscalClientConfig {
	return &InfiscalClient{
		AuthMode:        authMode,
		Credentials:     credentials,
		WorkspaceConfig: workspaceConfig,
	}
}

func (c *InfiscalClient) GetAllSecrets() map[string]string {

	cachedSecrets := helpers.ReturnAll()

	if helpers.HasLenInArray(cachedSecrets) {
		return cachedSecrets
	}

	return map[string]string{}
}
func (c *InfiscalClient) GetSecret(name string, options string) SecretBundle {
	return SecretBundle{}
}
func (c *InfiscalClient) CreateSecret(name, value string) SecretBundle {
	return SecretBundle{}
}
func (c *InfiscalClient) DeleteSecret(name string) {

}
