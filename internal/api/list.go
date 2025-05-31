package api

import (
	"encoding/json"
	"net/http"

	"github.com/mistcommunity/cloud-vm-api/internal/auth"
	"github.com/mistcommunity/cloud-vm-api/internal/cloud"
	"github.com/mistcommunity/cloud-vm-api/internal/config"
)

type VMListResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	// Add more fields if you want, such as IPs, Region, etc.
}

func ListVMsHandler(w http.ResponseWriter, r *http.Request) {
	// Parse environment param (could also be a path param or in query)
	envName := r.URL.Query().Get("environment")
	if envName == "" {
		http.Error(w, "Missing environment param", http.StatusBadRequest)
		return
	}

	env, err := config.GetEnvironment(envName)
	if err != nil {
		http.Error(w, "Environment not found", http.StatusNotFound)
		return
	}

	// Get bearer token for credentials
	token := r.Header.Get("Authorization")
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	} else {
		http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
		return
	}

	creds, err := auth.DecodeCredentials(token)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Get cloud provider
	provider, err := cloud.NewCloudProvider(env, creds)
	if err != nil {
		http.Error(w, "Failed to initialize cloud provider: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Get region - usually you need region ID from env or config
	regionId := env.Region
	region, err := provider.GetIRegionById(regionId)
	if err != nil {
		http.Error(w, "Failed to get region: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Get host (usually the first host)
	hosts, err := region.GetIHosts()
	if err != nil || len(hosts) == 0 {
		http.Error(w, "No hosts found in region", http.StatusInternalServerError)
		return
	}
	host := hosts[0] // Or select by your own criteria

	// List VMs
	vms, err := host.GetIVMs()
	if err != nil {
		http.Error(w, "Failed to list VMs: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := []VMListResponse{}
	for _, vm := range vms {
		resp = append(resp, VMListResponse{
			ID:     vm.GetGlobalId(),
			Name:   vm.GetName(),
			Status: vm.GetStatus(),
			// Add more as needed
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
