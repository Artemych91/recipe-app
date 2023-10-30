package models

type Recipe struct {
	ID          int    `json:"id"`
	UserID      int    `json:"userId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
}
