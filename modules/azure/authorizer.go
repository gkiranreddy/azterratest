package azure

import (
	"os"

	"github.com/Azure/go-autorest/autorest"
	az "github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

const (
	// AuthFromEnvClient is an env variable supported by the Azure SDK
	AuthFromEnvClient = "AZURE_CLIENT_ID"

	// AuthFromEnvSecret
	AuthFromEnvSecret = "AZURE_CLIENT_SECRET"

	// AuthFromEnvSubscription
	AuthFromEnvSubscription = "AZURE_SUBSCRIPTION_ID"

	// AuthFromEnvTenant is an env variable supported by the Azure SDK
	AuthFromEnvTenant = "AZURE_TENANT_ID"

	// AuthFromFile is an env variable supported by the Azure SDK
	AuthFromFile = "AZURE_AUTH_LOCATION"
	
)

// NewAuthorizer creates an Azure authorizer adhering to standard auth mechanisms provided by the Azure Go SDK
// See Azure Go Auth docs here: https://docs.microsoft.com/en-us/go/azure/azure-sdk-go-authorization
func NewAuthorizer() (*autorest.Authorizer, error) {
	// Carry out env var lookups
	_, clientIDExists := os.LookupEnv(AuthFromEnvClient)
	_, clientSecretExists := os.LookupEnv(AuthFromEnvSecret)
	_, clientSubExists := os.LookupEnv(AuthFromEnvSubscription)
	_, tenantIDExists := os.LookupEnv(AuthFromEnvTenant)
	_, fileAuthSet := os.LookupEnv(AuthFromFile)

	// Execute logic to return an authorizer from the correct method
	if clientIDExists && tenantIDExists && clientSecretExists && clientSubExists {
		authorizer, err := auth.NewAuthorizerFromEnvironment()
		return &authorizer, err
	} else if fileAuthSet {
		authorizer, err := auth.NewAuthorizerFromFile(az.PublicCloud.ResourceManagerEndpoint)
		return &authorizer, err
	} else {
		authorizer, err := auth.NewAuthorizerFromCLI()
		return &authorizer, err
	}
}
