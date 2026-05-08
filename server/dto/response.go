package dto

type AuthResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreatePaymentResponse struct {
	SnapToken   string `json:"snap_token"`
	RedirectURL string `json:"redirect_url"`
}