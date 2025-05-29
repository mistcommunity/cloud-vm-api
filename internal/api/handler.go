package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mistcommunity/cloud-vm-api/internal/auth"
	"github.com/mistcommunity/cloud-vm-api/internal/cloud"
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

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/vm", CreateVMHandler).Methods("POST")
	return r
}

func CreateVMHandler(w http.ResponseWriter, r *http.Request) {
	var req VMCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Get environment config
	env, err := config.GetEnvironment(req.Environment)
	if err != nil {
		http.Error(w, "Environment not found", http.StatusNotFound)
		return
	}

	// Get bearer token
	token := r.Header.Get("Authorization")
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	} else {
		http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
		return
	}

	// Decode token
	creds, err := auth.DecodeCredentials(token)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Look up machine type config
	mtConfig, err := config.GetMachineTypeConfig(req.MachineType, env.Cloud)
	if err != nil {
		http.Error(w, "Invalid machinetype for this cloud", http.StatusBadRequest)
		return
	}

	// Get host (from cloudmux)
	host, err := cloud.NewCloudProvider(env, creds)
	if err != nil {
		http.Error(w, "Failed to initialize cloud provider: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Compose VM config for cloudmux
	vmConfig := &cloudprovider.SManagedVMCreateConfig{
		Name:              req.Name,
		ExternalImageId:   mtConfig.Image,
		InstanceType:      mtConfig.InstanceType,
		ExternalNetworkId: env.NetworkId,
		ExternalVpcId:     env.Subnet,
		UserData:          req.CloudInit,
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
