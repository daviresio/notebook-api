package model

type Tag struct {
	BaseModel
	Name string `gorm:"not null" json:"name"`
	Notebooks []Notebook `gorm:"many2many:notebook_tags;" json:"notebooks,omitempty"`
}
