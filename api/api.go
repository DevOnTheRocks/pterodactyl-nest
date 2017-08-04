package api

import "github.com/cossacksman/nest/network"
import "encoding/json"
import "fmt"

func getResponse(endpoint string) []byte {
	headers := network.GetAuthHeaders(endpoint)
	return network.Get(endpoint, nil, headers)
}

func putResponse(endpoint string, body interface{}) []byte {
	headers := network.GetAuthHeaders(endpoint)
	return network.Put(endpoint, body, headers)
}

// GetServers - Return Pterodactyl Servers API response.
func GetServers() ServersResponse {
	var servers ServersResponse
	response := getResponse("admin/servers")
	json.Unmarshal(response, &servers)

	return servers
}

// GetServer - Return Pterodactyl Server API response.
func GetServer(id int) ServerResponse {
	var server ServerResponse
	response := getResponse(fmt.Sprintf("admin/servers/%d", id))
	json.Unmarshal(response, &server)

	return server
}

// GetUsers - Return Pterodactyl Users API response.
func GetUsers() UsersResponse {
	var users UsersResponse
	response := getResponse("admin/users")
	json.Unmarshal(response, &users)

	return users
}

// GetUser - Return Pterodactyl User API response.
func GetUser(id int) UserResponse {
	var user UserResponse
	response := getResponse(fmt.Sprintf("admin/users/%d", id))
	json.Unmarshal(response, &user)

	return user
}

// GetNodes - Return Pterodactyl Users API response.
func GetNodes() NodesResponse {
	var nodes NodesResponse
	response := getResponse("admin/nodes")
	json.Unmarshal(response, &nodes)

	return nodes
}

// GetNode - Return Pterodactyl Node API response.
func GetNode(id int) NodeResponse {
	var node NodeResponse
	response := getResponse(fmt.Sprintf("admin/nodes/%d", id))
	json.Unmarshal(response, &node)

	return node
}

// TogglePower - Return Pterodactyl Power API response.
func TogglePower(id int, action string) {
	server := GetServer(id).Data.Attributes

	body := struct {
		Action string `json:"action"`
	}{action}

	response := putResponse(fmt.Sprintf("me/server/%s", server.UUIDShort), body)

	fmt.Printf("%+v", response)
}
