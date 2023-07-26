package types

type Post struct {
	ID          int `json:"ID"`
	Title       int `json:"Title"`
	Description int `json:"Description"`
	AuthorID    int `json:"AuthorId"`
}
