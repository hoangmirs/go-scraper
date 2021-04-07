package models

type OAuthError struct {
	Code    string `json:"error"`
	Message string `json:"error_description"`
}
