package models

type Item struct {
	Model
	Title       string `gorm:"title" json:"title"`
	Description string `gorm:"description" json:"description"`
	IsFinish    bool   `gorm:"is_finish" json:"is_finish`
}
