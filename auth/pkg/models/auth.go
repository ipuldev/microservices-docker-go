package models

// Catalog is used to represent Catalog profile data
type Auth struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type HttpResponse struct {
	Message   string    `json:"message"`
	OauthData Oauth2Key `json:"oauth_data"`
}

type Oauth2Key struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}
