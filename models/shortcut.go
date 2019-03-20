package models

import (
	"github.com/grt1st/short_links/core"
	"github.com/jinzhu/gorm"
	"math/rand"
	"time"
)

type Shortcut struct {
	gorm.Model
	Short  string `gorm:"not null;unique"`
	Value  string `gorm:"not null;"`
	Source string `gorm:"default:'web'"`
}

func InitSql() {
	core.DB.AutoMigrate(&Shortcut{})

}

func NewShortcut(value string) (*Shortcut, error) {
	s := &Shortcut{
		Short: genRandomString(6),
		Value: value,
	}
	err := core.DB.Create(s).Error
	return s, err
}

func GetShortcutByShort(value string) Shortcut {
	var s Shortcut
	core.DB.Where("short = ?", value).First(&s)
	return s
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func genRandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
