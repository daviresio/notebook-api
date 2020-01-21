package model

type Address struct {
	BaseModel
	Address string `gorm:"not null" json:"address"`
	City string `gorm:"not null" json:"city"`
	State string `gorm:"not null" json:"state"`
	PostalCode string `gorm:"not null" json:"postal_code"`
}
