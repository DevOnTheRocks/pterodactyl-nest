package main

// ServersResponse - Pterodactyl API response.
type ServersResponse struct {
	Data []struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			UUID         string      `json:"uuid"`
			UUIDShort    string      `json:"uuidShort"`
			NodeID       int         `json:"node_id"`
			Name         string      `json:"name"`
			Description  string      `json:"description"`
			SkipScripts  bool        `json:"skip_scripts"`
			Suspended    int         `json:"suspended"`
			OwnerID      int         `json:"owner_id"`
			Memory       int         `json:"memory"`
			Swap         int         `json:"swap"`
			Disk         int         `json:"disk"`
			Io           int         `json:"io"`
			CPU          int         `json:"cpu"`
			OomDisabled  int         `json:"oom_disabled"`
			AllocationID int         `json:"allocation_id"`
			ServiceID    int         `json:"service_id"`
			OptionID     int         `json:"option_id"`
			PackID       interface{} `json:"pack_id"`
			Startup      string      `json:"startup"`
			Image        string      `json:"image"`
			Username     string      `json:"username"`
			Installed    int         `json:"installed"`
			CreatedAt    string      `json:"created_at"`
			UpdatedAt    string      `json:"updated_at"`
		} `json:"attributes"`
	} `json:"data"`
	Meta struct {
		Pagination struct {
			Total       int `json:"total"`
			Count       int `json:"count"`
			PerPage     int `json:"per_page"`
			CurrentPage int `json:"current_page"`
			TotalPages  int `json:"total_pages"`
		} `json:"pagination"`
	} `json:"meta"`
	Links struct {
		Self  string `json:"self"`
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"links"`
}

// ServerResponse - Pterodactyl single server API response.
type ServerResponse struct {
	Data struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			UUID         string      `json:"uuid"`
			UUIDShort    string      `json:"uuidShort"`
			NodeID       int         `json:"node_id"`
			Name         string      `json:"name"`
			Description  string      `json:"description"`
			SkipScripts  bool        `json:"skip_scripts"`
			Suspended    int         `json:"suspended"`
			OwnerID      int         `json:"owner_id"`
			Memory       int         `json:"memory"`
			Swap         int         `json:"swap"`
			Disk         int         `json:"disk"`
			Io           int         `json:"io"`
			CPU          int         `json:"cpu"`
			OomDisabled  int         `json:"oom_disabled"`
			AllocationID int         `json:"allocation_id"`
			ServiceID    int         `json:"service_id"`
			OptionID     int         `json:"option_id"`
			PackID       interface{} `json:"pack_id"`
			Startup      string      `json:"startup"`
			Image        string      `json:"image"`
			Username     string      `json:"username"`
			Installed    int         `json:"installed"`
			CreatedAt    string      `json:"created_at"`
			UpdatedAt    string      `json:"updated_at"`
		} `json:"attributes"`
	} `json:"data"`
}
