package domain

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"log"
	"time"
)

type Base struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key"`
	CreatedAt time.Time `json:"createdAT" gorm:"type:date"`
	UpdatedAt time.Time `json:"updatedAT" gorm:"type:date"`
	DeletedAt time.Time `json:"deletedAt" gorm:"type:date" sql:"index"`
}

func NewUser() *User {
	return &User{}
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("CreatedAt", time.Now())

	if err != nil {
		log.Fatalf("Error obj creating: %v", err)
	}

	err = scope.SetColumn("ID", uuid.NewV4().String())

	if err != nil {
		log.Fatalf("Error obj creating: %v", err)
	}

	return nil
}
