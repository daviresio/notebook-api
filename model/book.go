package model

type Book struct {
	BaseModel
	Name string `gorm:"not null" json:"name"`
	Notebooks []Notebook `gorm:"foreignkey:BookId" json:"notebooks"`
}
