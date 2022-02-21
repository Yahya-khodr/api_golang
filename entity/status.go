package entity

type Status struct {
	ID      int    `json:"post_id"`
	Content string `json:"post_content"`
	UserID  int    `json:"created_by"`
}
