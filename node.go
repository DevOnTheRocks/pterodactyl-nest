package main

// NodesResponse - A nodes Pterodactyl API response.
type NodesResponse []struct {
	ID                 int    `json:"id"`
	Public             int    `json:"public"`
	Name               string `json:"name"`
	LocationID         int    `json:"location_id"`
	Fqdn               string `json:"fqdn"`
	Scheme             string `json:"scheme"`
	Memory             int    `json:"memory"`
	MemoryOverallocate int    `json:"memory_overallocate"`
	Disk               int    `json:"disk"`
	DiskOverallocate   int    `json:"disk_overallocate"`
	UploadSize         int    `json:"upload_size"`
	DaemonListen       int    `json:"daemonListen"`
	DaemonSFTP         int    `json:"daemonSFTP"`
	DaemonBase         string `json:"daemonBase"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
}

// NodeResponse - A single node Pterodactyl response.
type NodeResponse struct {
	Node struct {
		ID                 int    `json:"id"`
		Public             int    `json:"public"`
		Name               string `json:"name"`
		LocationID         int    `json:"location_id"`
		Fqdn               string `json:"fqdn"`
		Scheme             string `json:"scheme"`
		Memory             int    `json:"memory"`
		MemoryOverallocate int    `json:"memory_overallocate"`
		Disk               int    `json:"disk"`
		DiskOverallocate   int    `json:"disk_overallocate"`
		UploadSize         int    `json:"upload_size"`
		DaemonListen       int    `json:"daemonListen"`
		DaemonSFTP         int    `json:"daemonSFTP"`
		DaemonBase         string `json:"daemonBase"`
		CreatedAt          string `json:"created_at"`
		UpdatedAt          string `json:"updated_at"`
		Allocations        []struct {
			ID       int         `json:"id"`
			IP       string      `json:"ip"`
			IPAlias  interface{} `json:"ip_alias"`
			Port     int         `json:"port"`
			ServerID interface{} `json:"server_id"`
		} `json:"allocations"`
	} `json:"node"`
}
