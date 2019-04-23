package model

//Book sd
type Book struct {
	ID     string `json:"_id"`
	Title  string `json:"title"`
	Author Author `json:"author"`
}
