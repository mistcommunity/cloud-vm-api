package aws

import (
	"github.com/mistcommunity/cloud-vm-api/internal/auth"
	"github.com/mistcommunity/cloud-vm-api/internal/config"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud/aws"
	"yunion.io/x/cloudmux/pkg/multicloud/aws/provider"
)

type AWSProvider struct {
	client *aws.SAwsClient
	Host   *aws.SHost
}

func NewAWSProvider(env *config.Environment, creds *auth.Credentials) (*AWSProvider, error) {
	cfg := cloudprovider.ProviderConfig{
		Id:      "1",
		Name:    "aws-1",
		Vendor:  "AWS",
		URL:     "InternationalCloud",
		Account: creds.AccessKey,
		Secret:  creds.Secret,
	}
	p, err := provider.GetProvider(cfg)
	if err != nil {
		return nil, err
	}
	return p, nil
	/*
			cfg := aws.NewAwsClientConfig(
				"InternationalCloud",
				creds.AccessKey,
				creds.Secret,
				env.AccountId,
			)

			cfg.CloudproviderConfig(providerCfg)

			cfg.Debug(true)


		client, err := aws.NewAwsClient(cfg)
		if err != nil {
			return nil, err
		}
	*/

	/*
		region, err := client.GetRegion(env.Region)
		if err != nil {
			return nil, err
		}
		zone, err := region.GetIZoneById()
		if err != nil {
			return nil, err
		}
		iwires := zone[0].GetIWires()
		if len(iwires) == 0 {
			return nil, errors.New("no iwires found")
		}
	*/

	/*
		host := &aws.SHost{}
		return &AWSProvider{
			client: client,
			Host:   host,
		}, nil
	*/
}
