package api

import "github.com/cossacksman/nest/network"
import "encoding/json"

// GetServers - Return Pterodactyl Servers API response.
func GetServers() ServersResponse {
	endpoint := "admin/servers"
	headers := network.GetAuthHeaders(endpoint)
	response := network.Get(endpoint, nil, headers)

	var servers ServersResponse
	json.Unmarshal(response, &servers)

	return servers
}
