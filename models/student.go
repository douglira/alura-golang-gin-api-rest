package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name                      string `json:"name,omitempty" validate:"nonzero"`
	IdentityNumber            string `json:"identityNumber,omitempty" gorm:"unique" validate:"nonzero,regexp=^[0-9]*$,iscpf"`
	GeneralRegistrationNumber string `json:"generalRegistrationNumber,omitempty" validate:"len=9,regexp=^[0-9]*$"`
}

func StudentValidator(student *Student) error {
	if err := validator.Validate(student); err != nil {
		return err
	}
	return nil
}
