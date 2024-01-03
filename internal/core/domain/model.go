package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RegularText struct {
	UUID    string `gorm:"primaryKey"`
	Content string `gorm:"not null"`
}

func (t *RegularText) BeforeCreate(_ *gorm.DB) (err error) {
	t.UUID = uuid.NewString()
	return
}

type CodeExample struct {
	UUID                    string `gorm:"primaryKey"`
	Content                 string `gorm:"not null"`
	ProgrammingLanguageUUID string `gorm:"not null;index"`
	ProgrammingLanguage     ProgrammingLanguage
}

func (e *CodeExample) BeforeCreate(_ *gorm.DB) (err error) {
	e.UUID = uuid.NewString()
	return
}

type ProgrammingLanguage struct {
	UUID string `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}

func (l *ProgrammingLanguage) BeforeCreate(_ *gorm.DB) (err error) {
	l.UUID = uuid.NewString()
	return
}