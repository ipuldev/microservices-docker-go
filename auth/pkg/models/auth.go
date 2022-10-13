package models

// Catalog is used to represent Catalog profile data
type Auth struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type HttpResponse struct {
	Message string `json:"message"`
}
