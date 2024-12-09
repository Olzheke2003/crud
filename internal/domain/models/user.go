package models

type User struct {
	ID string `json:"id"`
	UserProperties
	DateProperties
}

type UserProperties struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type DateProperties struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
