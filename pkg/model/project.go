package model

type Project struct {
	ID          *string `json:"id " db:"id"`
	Name        *string `json:"name,omitempty" db:"name"`
	Description *string `json:"description,omitempty" db:"description"`
	Picture     *string `json:"picture,omitempty" db:"picture"`
}
