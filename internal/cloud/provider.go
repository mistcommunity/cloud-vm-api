package cloud

import (
	"github.com/mistcommunity/cloud-vm-api/internal/auth"
	"github.com/mistcommunity/cloud-vm-api/internal/config"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	_ "yunion.io/x/cloudmux/pkg/multicloud/aws/provider"
	_ "yunion.io/x/cloudmux/pkg/multicloud/azure/provider"
)

func NewCloudProvider(env *config.Environment, creds *auth.Credentials) (cloudprovider.ICloudProvider, error) {
	// cfg.Vendor must be exactly "AWS", "Azure", etc. (case-sensitive)
	cfg := cloudprovider.ProviderConfig{
		Id:      "1",
		Name:    "aws-1",
		Vendor:  env.Cloud,
		URL:     env.Url,
		Account: creds.AccessKey,
		Secret:  creds.Secret,
	}
	return cloudprovider.GetProvider(cfg)
}
