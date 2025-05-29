package azure

import (
	"github.com/pkg/errors"
	"yunion.io/x/cloudmux/pkg/multicloud/azure"

	"github.com/mistcommunity/cloud-vm-api/internal/auth"
	"github.com/mistcommunity/cloud-vm-api/internal/config"
)

type AzureProvider struct {
	client *azure.SAzureClient
	Host   *azure.SHost
}

func NewAzureProvider(env *config.Environment, creds *auth.Credentials) (*AzureProvider, error) {
	cfg := azure.NewAzureClientConfig(
		"AzurePublicCloud",
		"your-tenant-id",
		creds.AccessKey,
		creds.Secret,
	)
	cfg.SubscriptionId(env.SubscriptionId)
	cfg.Debug(true)

	client, err := azure.NewAzureClient(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Azure client")
	}

	/*
		region, err := client.GetRegion(env.Region)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get Azure region")
		}

		zone, err := region.GetZone("")
		if err != nil {
			return nil, errors.Wrap(err, "failed to get Azure zone")
		}
	*/

	return &AzureProvider{
		client: client,
		Host:   &azure.SHost{},
	}, nil
}
