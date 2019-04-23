package model

//Book sd
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author Author `json:"author"`
}
