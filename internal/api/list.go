package api

import (
	"encoding/json"
	"net/http"
)

type VMListResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	// Add more fields if you want, such as IPs, Region, etc.
}

func ListVMsHandler(w http.ResponseWriter, r *http.Request) {
	host, _, err, status := getHost(r)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

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
