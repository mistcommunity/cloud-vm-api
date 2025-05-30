package nutanix

import (
	"github.com/pkg/errors"
	"yunion.io/x/cloudmux/pkg/multicloud/nutanix"

	"github.com/mistcommunity/cloud-vm-api/internal/auth"
	"github.com/mistcommunity/cloud-vm-api/internal/config"
)

type NutanixProvider struct {
	client *nutanix.SNutanixClient
	Host   *nutanix.SHost
}

func NewNutanixProvider(env *config.Environment, creds *auth.Credentials) (*NutanixProvider, error) {
	cfg := nutanix.NewNutanixClientConfig(
		env.Host,
		creds.AccessKey,
		creds.Secret,
		9440,
	)
	cfg.Debug(true)

	client, err := nutanix.NewNutanixClient(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Nutanix client")
	}

	/*
			region, err := client.GetRegion("")
			if err != nil {
				return nil, errors.Wrap(err, "failed to get Nutanix region")
			}

			zone, err := region.GetZone("")
			if err != nil {
				return nil, errors.Wrap(err, "failed to get Nutanix zone")
			}


		hosts, err := region.GetHosts()
		if err != nil {
			return nil, errors.Wrap(err, "failed to get Nutanix hosts")
		}
		if len(hosts) == 0 {
			return nil, errors.New("no hosts found")
		}
	*/

	// &hosts[0]
	return &NutanixProvider{
		client: client,
		Host:   nil,
	}, nil
}
