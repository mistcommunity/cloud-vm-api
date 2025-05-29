package aliyun

import (
	"github.com/pkg/errors"
	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"

	"github.com/mistcommunity/cloud-vm-api/internal/auth"
	"github.com/mistcommunity/cloud-vm-api/internal/config"
)

type AliyunProvider struct {
	client *aliyun.SAliyunClient
	Host   *aliyun.SHost
}

func NewAliyunProvider(env *config.Environment, creds *auth.Credentials) (*AliyunProvider, error) {
	cfg := aliyun.NewAliyunClientConfig(
		"InternationalCloud",
		creds.AccessKey,
		creds.Secret,
	)
	cfg.Debug(true)

	client, err := aliyun.NewAliyunClient(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Aliyun client")
	}

	/*
		region := client.GetRegion(env.Region)
		zones, err := region.GetClient().GetAllZones()
		if err != nil {
			return nil, errors.Wrap(err, "failed to get Aliyun zones")
		}

		// FixMe
		zone := zones[0]
	*/

	return &AliyunProvider{
		client: client,
		Host:   &aliyun.SHost{},
	}, nil
}
