package cloud

import (
	"github.com/mistcommunity/cloud-vm-api/internal/auth"
	"github.com/mistcommunity/cloud-vm-api/internal/cloud/aws"
	"github.com/mistcommunity/cloud-vm-api/internal/config"
	"github.com/pkg/errors"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

func NewCloudProvider(env *config.Environment, creds *auth.Credentials) (cloudprovider.ICloudProvider, error) {
	switch env.Cloud {
	case "aws":
		prov, err := aws.NewAWSProvider(env, creds)
		if err != nil {
			return nil, err
		}
		return prov, nil
	default:
		return nil, errors.Errorf("unsupported cloud provider: %s", env.Cloud)
	}
}
