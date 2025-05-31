# Cloud API Server

This is a REST API server for creating virtual machines across multiple cloud providers using a unified interface.

## Prerequisites

- Go 1.16+
- Valid credentials for the target cloud providers
- Configured environments.json file

## Setup

1. Ensure environments.json is configured with your cloud environments
2. Run the server:
```bash
go run cmd/main.go
```

## Usage
Bearer token for all platforms are created with command like:
```bash
echo '{"access_key":"<access key / username>","secret":"<secret / password>"}' | base64 -w 0
```

### List VMs
```bash
curl -X GET http://localhost:8080/vm?environment=aws-test1 \
  -H "Authorization: Bearer <Bearer token>"
```

### Create VM
Create a VM by sending a POST request to /vm:
```bash
cloud_init=$(cat <<EOF | base64
#cloud-config
users:
- name: dev
  sudo: ALL=(ALL) NOPASSWD:ALL
EOF
)
echo $cloud_init

curl -X POST http://localhost:8080/vm \
  -H "Authorization: Bearer <Bearer token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "test1",
    "environment": "<env name>",
    "machinetype": "medium-debian",
    "cloud_init": "<base64 encoded cloud init>"
  }'
```

Credentials should be base64-encoded JSON with the following format:
```json
{
  "access_key": "your-access-key",
  "secret": "your-secret-key"
}
```

## Notes

- The API is stateless and handles all backend calls synchronously
- Authentication is handled via Bearer token containing base64-encoded credentials
- The server supports Alibaba Cloud, AWS, Azure, ESXi, Nutanix, and Proxmox
- Additional cloud providers can be added by implementing the CloudProvider interface
