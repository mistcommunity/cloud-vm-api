package aws

import (
	"github.com/mistcommunity/cloud-vm-api/internal/auth"
	"github.com/mistcommunity/cloud-vm-api/internal/config"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	_ "yunion.io/x/cloudmux/pkg/multicloud/aws/provider"
)

func NewAWSProvider(env *config.Environment, creds *auth.Credentials) (cloudprovider.ICloudProvider, error) {
	cfg := cloudprovider.ProviderConfig{
		Id:      "1",
		Name:    "aws-1",
		Vendor:  "AWS",
		URL:     "InternationalCloud",
		Account: creds.AccessKey,
		Secret:  creds.Secret,
	}
	factory, err := cloudprovider.GetProviderFactory("AWS")
	if err != nil {
		return nil, err
	}
	provider, err := factory.GetProvider(cfg)
	if err != nil {
		return nil, err
	}
	return provider, nil
}
