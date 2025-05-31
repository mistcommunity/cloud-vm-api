package api

import (
	"encoding/json"
	"net/http"

	"github.com/mistcommunity/cloud-vm-api/internal/config"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type VMCreateRequest struct {
	Name        string `json:"name"`
	Environment string `json:"environment"`
	MachineType string `json:"machinetype"`
	CloudInit   string `json:"cloud_init,omitempty"`
}

type VMCreateResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func CreateVMHandler(w http.ResponseWriter, r *http.Request) {
	host, env, err, status := getHost(r)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	var req VMCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Look up machine type config
	mtConfig, err := config.GetMachineTypeConfig(req.MachineType, env.Cloud)
	if err != nil {
		http.Error(w, "Invalid machinetype for this cloud: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Compose VM config for cloudmux
	vmConfig := &cloudprovider.SManagedVMCreateConfig{
		Name:              req.Name,
		ExternalImageId:   mtConfig.Image,
		InstanceType:      mtConfig.InstanceType,
		ExternalNetworkId: env.NetworkId,
		ExternalVpcId:     env.VpcId,
		UserData:          req.CloudInit,
		SysDisk: cloudprovider.SDiskInfo{
			StorageExternalId: env.StorageExternalId,
			SizeGB:            10,
		},
		Cpu:            2,
		MemoryMB:       2048,
		OsDistribution: "debian",
		// Fill more fields if you want, eg. SysDisk, Cpu, MemoryMB...
	}

	vm, err := host.CreateVM(vmConfig)
	if err != nil {
		http.Error(w, "Create VM failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := VMCreateResponse{
		ID:     vm.GetGlobalId(),
		Name:   vm.GetName(),
		Status: vm.GetStatus(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
