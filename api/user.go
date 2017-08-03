package api

// UsersResponse - Pterodactyl Users API response.
type UsersResponse struct {
	Data []struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			UUID      string `json:"uuid"`
			Username  string `json:"username"`
			Email     string `json:"email"`
			NameFirst string `json:"name_first"`
			NameLast  string `json:"name_last"`
			Language  string `json:"language"`
			RootAdmin int    `json:"root_admin"`
			UseTotp   int    `json:"use_totp"`
			Gravatar  int    `json:"gravatar"`
			CreatedAt string `json:"created_at"`
			UpdatedAt string `json:"updated_at"`
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

// UserResponse - Pterodactyl single user API response.
type UserResponse struct {
	Data struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			UUID      string `json:"uuid"`
			Username  string `json:"username"`
			Email     string `json:"email"`
			NameFirst string `json:"name_first"`
			NameLast  string `json:"name_last"`
			Language  string `json:"language"`
			RootAdmin int    `json:"root_admin"`
			UseTotp   int    `json:"use_totp"`
			Gravatar  int    `json:"gravatar"`
			CreatedAt string `json:"created_at"`
			UpdatedAt string `json:"updated_at"`
		} `json:"attributes"`
	} `json:"data"`
}
