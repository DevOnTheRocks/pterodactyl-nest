package api

import "github.com/cossacksman/nest/network"
import "encoding/json"
import "fmt"

// GetServers - Return Pterodactyl Servers API response.
func GetServers() ServersResponse {
	var servers ServersResponse
	response := network.Get("admin/servers")
	json.Unmarshal(response, &servers)

	return servers
}

// GetServer - Return Pterodactyl Server API response.
func GetServer(id int) ServerResponse {
	var server ServerResponse
	response := network.Get(fmt.Sprintf("admin/servers/%d", id))
	json.Unmarshal(response, &server)

	return server
}

// GetUsers - Return Pterodactyl Users API response.
func GetUsers() UsersResponse {
	var users UsersResponse
	response := network.Get("admin/users")
	json.Unmarshal(response, &users)

	return users
}

// GetUser - Return Pterodactyl User API response.
func GetUser(id int) UserResponse {
	var user UserResponse
	response := network.Get(fmt.Sprintf("admin/users/%d", id))
	json.Unmarshal(response, &user)

	return user
}

// GetNodes - Return Pterodactyl Users API response.
func GetNodes() NodesResponse {
	var nodes NodesResponse
	response := network.Get("admin/nodes")
	json.Unmarshal(response, &nodes)

	return nodes
}

// GetNode - Return Pterodactyl Node API response.
func GetNode(id int) NodeResponse {
	var node NodeResponse
	response := network.Get(fmt.Sprintf("admin/nodes/%d", id))
	json.Unmarshal(response, &node)

	return node
}

// TogglePower - Return Pterodactyl Power API response.
func TogglePower(id int, action string) string {
	server := GetServer(id).Data.Attributes

	data := struct {
		Action string `json:"action"`
	}{action}

	body, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	response := network.Post(fmt.Sprintf("user/server/%s/power", server.UUIDShort), string(body))

	return string(response)
}
