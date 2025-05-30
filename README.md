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

Create a VM by sending a POST request to /vm:

### ProxmoxVE
```bash
curl -X POST http://localhost:8080/vm \
  -H "Authorization: Bearer eyJzZWNyZXQiOiJRd2VydHk3ISIsImFjY2Vzc19rZXkiOiJyb290QHBhbSJ9" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "test1",
    "environment": "pve-test1",
    "machinetype": "medium",
    "cloud_init": "#cloud-config\nusers:\n  - name: dev\n    sudo: ALL=(ALL) NOPASSWD:ALL"
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
