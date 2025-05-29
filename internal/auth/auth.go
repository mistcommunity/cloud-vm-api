package auth

import (
	"encoding/base64"
	"encoding/json"

	"github.com/pkg/errors"
)

type Credentials struct {
	AccessKey string `json:"access_key"`
	Secret    string `json:"secret"`
}

func DecodeCredentials(token string) (*Credentials, error) {
	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode token")
	}

	var creds Credentials
	if err := json.Unmarshal(data, &creds); err != nil {
		return nil, errors.Wrap(err, "failed to parse credentials")
	}

	return &creds, nil
}
