package api

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
	Data struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			Public             int    `json:"public"`
			Name               string `json:"name"`
			LocationID         int    `json:"location_id"`
			Fqdn               string `json:"fqdn"`
			Scheme             string `json:"scheme"`
			BehindProxy        bool   `json:"behind_proxy"`
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
		} `json:"attributes"`
	} `json:"data"`
}
