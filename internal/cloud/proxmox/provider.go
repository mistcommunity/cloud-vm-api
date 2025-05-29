package proxmox

import (
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
		"host",
		8080,
	)
	cfg.Debug(true)

	client, err := proxmox.NewProxmoxClient(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Proxmox client")
	}

	/*
		region, err := client.GetRegion("")
		if err != nil {
			return nil, errors.Wrap(err, "failed to get Proxmox region")
		}

		zone, err := region.GetZone("")
		if err != nil {
			return nil, errors.Wrap(err, "failed to get Proxmox zone")
		}

		hosts, err := region.GetHosts()
		if err != nil {
			return nil, errors.Wrap(err, "failed to get Proxmox hosts")
		}
		if len(hosts) == 0 {
			return nil, errors.New("no hosts found")
		}
	*/

	// &hosts[0]
	return &ProxmoxProvider{
		client: client,
		Host:   nil,
	}, nil
}
