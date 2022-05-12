package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name                      string `json:"name,omitempty"`
	IdentityNumber            string `json:"identityNumber,omitempty" gorm:"primaryKey,unique"`
	GeneralRegistrationNumber string `json:"generalRegistrationNumber,omitempty"`
}
