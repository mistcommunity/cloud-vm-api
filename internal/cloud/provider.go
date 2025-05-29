package cloud

import (
	"github.com/mistcommunity/cloud-vm-api/internal/auth"
	"github.com/mistcommunity/cloud-vm-api/internal/cloud/aliyun"
	"github.com/mistcommunity/cloud-vm-api/internal/cloud/aws"
	"github.com/mistcommunity/cloud-vm-api/internal/cloud/azure"
	"github.com/mistcommunity/cloud-vm-api/internal/cloud/nutanix"
	"github.com/mistcommunity/cloud-vm-api/internal/cloud/proxmox"
	"github.com/mistcommunity/cloud-vm-api/internal/config"
	"github.com/pkg/errors"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type CloudProvider interface {
	GetHost() cloudprovider.ICloudHost
}

func NewCloudProvider(env *config.Environment, creds *auth.Credentials) (cloudprovider.ICloudHost, error) {
	switch env.Cloud {
	case "aliyun":
		prov, err := aliyun.NewAliyunProvider(env, creds)
		if err != nil {
			return nil, err
		}
		return prov.Host, nil
	case "aws":
		prov, err := aws.NewAWSProvider(env, creds)
		if err != nil {
			return nil, err
		}
		return prov.Host, nil
	case "azure":
		prov, err := azure.NewAzureProvider(env, creds)
		if err != nil {
			return nil, err
		}
		return prov.Host, nil
	// case "esxi":
	// 	prov, err := azure.NewAzureProvider(env, creds)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return prov.Host, nil
	case "nutanix":
		prov, err := nutanix.NewNutanixProvider(env, creds)
		if err != nil {
			return nil, err
		}
		return prov.Host, nil
	case "proxmox":
		prov, err := proxmox.NewProxmoxProvider(env, creds)
		if err != nil {
			return nil, err
		}
		return prov.Host, nil
	default:
		return nil, errors.Errorf("unsupported cloud provider: %s", env.Cloud)
	}
}
