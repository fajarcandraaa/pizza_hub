package presentation

type (
	OrderRequest struct {
		MenuCode      string `json:"menu_code"`
		CheffInitials string `json:"initials"`
	}
)
