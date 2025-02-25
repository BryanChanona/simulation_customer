package models



type User struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ResponseUSer struct {
    User []User `json:"users"`
}