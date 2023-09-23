package presentation

// Request
type (
	NewChefRequest struct {
		CheffInitials string `json:"initials"`
		CheffName     string `json:"name"`
		Username      string `json:"username"`
		Password      string `json:"password"`
	}
)

// Response
type (
	ChefResponse struct {
		ID            string `json:"id"`
		CheffInitials string `json:"initials"`
		CheffName     string `json:"name"`
		Username      string `json:"username"`
	}
)
