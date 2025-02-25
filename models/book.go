package models


type Book struct {
	ID     int    `json:"id,omitempty"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type ResponseBook struct{
	Book []Book `json:"books"`
}