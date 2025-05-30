package proxmox

import (
	"fmt"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud/proxmox"

	"github.com/mistcommunity/cloud-vm-api/internal/auth"
	"github.com/mistcommunity/cloud-vm-api/internal/config"
	"github.com/pkg/errors"
)

type ProxmoxProvider struct {
	client *proxmox.SProxmoxClient
	Host   *proxmox.SHost
}

func NewProxmoxProvider(env *config.Environment, creds *auth.Credentials) (*ProxmoxProvider, error) {
	cfg := proxmox.NewProxmoxClientConfig(
		creds.AccessKey,
		creds.Secret,
		env.Host,
		8006,
	)

	providerCfg := cloudprovider.ProviderConfig{
		Name: "pve",
	}
	cfg.CloudproviderConfig(providerCfg)

	cfg.Debug(true)

	client, err := proxmox.NewProxmoxClient(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Proxmox client")
	}

	region := client.GetRegion()
	fmt.Printf("DEBUG: Found region: %v\r\n", region.GetName())
	zone, err := region.GetZone()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get Proxmox zone")
	}
	fmt.Printf("DEBUG: Found zone: %v\r\n", zone.GetName())

	hosts, err := region.GetHosts()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get Proxmox hosts")
	}
	if len(hosts) == 0 {
		return nil, errors.New("no hosts found")
	}

	// &hosts[0]

	// vmId := self.zone.region.GetClusterVmMaxId()

	return &ProxmoxProvider{
		client: client,
		Host:   &proxmox.SHost{},
	}, nil
}
