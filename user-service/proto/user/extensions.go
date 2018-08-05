package go_micro_srv_user

import (
	"github.com/satori/go.uuid"
	"github.com/jinzhu/gorm"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("Id", uuid.String())
}