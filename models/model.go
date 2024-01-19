package models

type User struct {
	ID    uint   `gorm:"primarykey autoincrement" json:"id"`
	Name  string `gorm:"not null" json:"name,omitempty"`
	Email string `gorm:"not null" json:"email,omitempty"`
}

func (User) TableName() string {
	return "users"
}
