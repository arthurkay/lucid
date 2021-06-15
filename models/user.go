package models

// User data structure for mapping user
// information in the database as well as
// json API endpoints with the provided
// json mappins
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
