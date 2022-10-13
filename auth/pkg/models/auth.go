package models

// Catalog is used to represent Catalog profile data
type Product struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
