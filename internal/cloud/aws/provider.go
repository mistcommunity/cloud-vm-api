package aws

import (
	"github.com/mistcommunity/cloud-vm-api/internal/auth"
	"github.com/mistcommunity/cloud-vm-api/internal/config"
	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

type AWSProvider struct {
	client *aws.SAwsClient
	Host   *aws.SHost
}

func NewAWSProvider(env *config.Environment, creds *auth.Credentials) (*AWSProvider, error) {
	cfg := aws.NewAwsClientConfig(
		"InternationalCloud",
		creds.AccessKey,
		creds.Secret,
		env.AccountId,
	)
	cfg.Debug(true)

	client, err := aws.NewAwsClient(cfg)
	if err != nil {
		return nil, err
	}

	/*
		region, err := client.GetRegion(env.Region)
		if err != nil {
			return nil, err
		}
		zone, err := region.GetZone("")
		if err != nil {
			return nil, err
		}
	*/

	host := &aws.SHost{}
	return &AWSProvider{
		client: client,
		Host:   host,
	}, nil
}
