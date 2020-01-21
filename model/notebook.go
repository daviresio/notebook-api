package model

import (
	"encoding/json"
)

type Notebook struct {
	BaseModel
	Name string `gorm:"not null" json:"name"`
	Book Book `gorm:"foreignkey:BookId" json:"book"`
	BookId uint `json:"book_id"`
	Tags []Tag `gorm:"many2many:notebook_tags;" json:"tags,omitempty"`
}

type notebookWihoutBook struct {
	BaseModel
	Name string `json:"name"`
	Tags []Tag `json:"tags,omitempty"`
}

func (notebook Notebook) MarshalJSON() ([]byte, error) {

	if notebook.Book.Id == 0 {
		type domain notebookWihoutBook
		return json.Marshal(domain{
			BaseModel: BaseModel{
				Id:        notebook.Id,
				CreatedAt: notebook.CreatedAt,
				UpdatedAt: notebook.CreatedAt,
				DeletedAt: notebook.DeletedAt,
			},
			Name:      notebook.Name,
			Tags:      notebook.Tags,
		})
	}

	type domain Notebook
	return json.Marshal(domain(notebook))
}