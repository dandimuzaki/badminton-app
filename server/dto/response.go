package dto

type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AuthResponse struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

type CreatePaymentResponse struct {
	SnapToken   string `json:"snap_token"`
	RedirectURL string `json:"redirect_url"`
}