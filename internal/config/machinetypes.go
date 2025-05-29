package config

import (
	"encoding/json"
	"os"

	"github.com/pkg/errors"
)

type MachineTypeCloudConfig struct {
	InstanceType string `json:"instance_type"`
	Image        string `json:"image"`
}
type MachineTypes map[string]map[string]MachineTypeCloudConfig

var MachineTypesConfig MachineTypes

func LoadMachineTypes(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return errors.Wrap(err, "failed to read machinetypes file")
	}
	return json.Unmarshal(data, &MachineTypesConfig)
}

func GetMachineTypeConfig(machineType, cloud string) (*MachineTypeCloudConfig, error) {
	if mt, ok := MachineTypesConfig[machineType]; ok {
		if conf, ok2 := mt[cloud]; ok2 {
			return &conf, nil
		}
		return nil, errors.Errorf("cloud %s not found for machinetype %s", cloud, machineType)
	}
	return nil, errors.Errorf("machinetype %s not found", machineType)
}

// Call from config/init()
func init() {
	_ = LoadMachineTypes("machinetypes.json")
}
