package models

// DB model
type Picture struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Author   string `json:"author"`
	Year     int    `json:"year"`
	Material string `json:"material"`
}

// Check data
type CreatePicture struct {
	Name     string `json:"name" binding:"required"`
	Author   string `json:"author" binding:"required"`
	Year     int    `json:"year" binding:"required"`
	Material string `json:"material" binding:"required"`
}

type UpdatePicture struct {
	Name     string `json:"name"`
	Author   string `json:"author"`
	Year     int    `json:"year"`
	Material string `json:"material"`
}
