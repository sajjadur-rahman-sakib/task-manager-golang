package models

type Task struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	UserID    uint
}
