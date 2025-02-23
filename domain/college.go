package domain

type College struct {
	ID     string `json:"id"`
	Name string `json:"name" binding:"required"`
	Location   string    `json:"location" binding:"required"`
	Duration string `json:"duration" binding:"required"`
	AccepRate string `json:"acceptRate" binding:"required"`
	NetPrice string `json:"netPrice" binding:"required"`
}
