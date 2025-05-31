package api

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

func StopVMHandler(w http.ResponseWriter, r *http.Request) {
	host, _, err, status := getHost(r)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	vars := mux.Vars(r)
	vmID := vars["id"]
	vm, err := getVMByID(host, vmID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	ctx := context.Background()
	if err := vm.StopVM(ctx, nil); err != nil {
		http.Error(w, "Failed to stop VM: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
