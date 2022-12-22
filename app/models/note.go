package models

type Note struct {
	ID      uint64 `json:"id"`
	UserId  uint64 `json:"user_id"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	Private bool   `json:"private"`
}
