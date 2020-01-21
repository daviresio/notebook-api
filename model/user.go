package model

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	Name string `json:"name"`
	Password string `gorm:"not null" json:"password"`
	Email string `gorm:"unique; not null" json:"email" binding:"required"`
	Address Address `gorm:"foreignkey:AddressId" json:"address"`
	AddressId uint `json:"address_id"`
}

func (user *User) BeforeCreate() (err error) {
	user.Password = EncryptPassword(user.Password)
	return nil
}

func EncryptPassword(password string) string {
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	return string(passwordHash)
}

func (user *User) AfterSave() (err error)  {
	user.RemovePassword()
	return nil
}

func (user *User) RemovePassword() {
	user.Password = ""
}

type userWihoutAddress struct {
	BaseModel
	Name string `json:"name"`
	Email string `json:"email"`
}

func (user User) MarshalJSON() ([]byte, error) {
	if user.Address.Id == 0 {
		type domain userWihoutAddress
		return json.Marshal(domain{
			BaseModel: BaseModel{
				Id:        user.Id,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.CreatedAt,
				DeletedAt: user.DeletedAt,
			},
			Name:      user.Name,
			Email:      user.Email,
		})
	}

	type domain User
	return json.Marshal(domain(user))
}