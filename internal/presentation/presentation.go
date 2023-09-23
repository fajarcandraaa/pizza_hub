package presentation

type LoginRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type MetaPagination struct {
	SortBy  string `json:"sortby,omitempty"`
	OrderBy string `json:"orderby,omitempty"`
	PerPage int    `json:"perpage,omitempty"`
	Page    int    `json:"page,omitempty"`
}
