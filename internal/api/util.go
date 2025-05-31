package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mistcommunity/cloud-vm-api/internal/auth"
	"github.com/mistcommunity/cloud-vm-api/internal/cloud"
	"github.com/mistcommunity/cloud-vm-api/internal/config"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

func getHost(r *http.Request) (cloudprovider.ICloudHost, *config.Environment, error, int) {
	envName := ""
	if r.Method == "POST" {
		var req VMCreateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return nil, nil, fmt.Errorf("Invalid request body"), http.StatusBadRequest
		}
		envName = r.URL.Query().Get("environment")
	} else {
		envName = r.URL.Query().Get("environment")
	}
	if envName == "" {
		return nil, nil, fmt.Errorf("Missing environment param"), http.StatusBadRequest
	}
	env, err := config.GetEnvironment(envName)
	if err != nil {
		return nil, nil, fmt.Errorf("Environment with name %s not found", envName), http.StatusBadRequest
	}

	token := r.Header.Get("Authorization")
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	} else {
		return nil, nil, fmt.Errorf("Missing or invalid Authorization header"), http.StatusUnauthorized
	}
	creds, err := auth.DecodeCredentials(token)
	if err != nil {
		return nil, nil, fmt.Errorf("Unable parsing credentials"), http.StatusUnauthorized
	}

	provider, err := cloud.NewCloudProvider(env, creds)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to initialize cloud provider: %s", err), http.StatusInternalServerError
	}
	regionId := env.Region
	region, err := provider.GetIRegionById(regionId)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to get region: %s", err), http.StatusInternalServerError
	}
	hosts, err := region.GetIHosts()
	if err != nil || len(hosts) == 0 {
		return nil, nil, fmt.Errorf("No hosts found in region: %s", env.Region), http.StatusInternalServerError
	}

	// FixMe: We might want to select other host too than just first one?
	host := hosts[0]
	return host, env, nil, http.StatusOK
}

func getVMByID(host cloudprovider.ICloudHost, vmID string) (cloudprovider.ICloudVM, error) {
	vms, err := host.GetIVMs()
	if err != nil {
		return nil, err
	}
	for _, vm := range vms {
		if vm.GetGlobalId() == vmID {
			return vm, nil
		}
	}
	return nil, fmt.Errorf("VM not found")
}
