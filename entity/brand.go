package entity

import (
	"errors"
	"fmt"
	"time"
)


type Brand struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Logo string `json:"logo"`
	Level string `json:"level"`
	IsActive bool `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

func (b *Brand) Validate() error {
	EnumBrandLevel := map[string]bool{
		"small": true,
		"medium": true,
		"large": true,
	}
	if b.Name == "" {
		return errors.New(fmt.Sprintf(ErrFieldRequired, "name"))
	}
	if b.Description == "" {
		return errors.New(fmt.Sprintf(ErrFieldRequired, "description"))
	}
	if b.Logo == "" {
		return errors.New(fmt.Sprintf(ErrFieldRequired, "logo"))
	}
	if b.Level == "" {
		return errors.New(fmt.Sprintf(ErrFieldRequired, "logo"))
	}
	if !EnumBrandLevel[b.Level] {
		return errors.New("level value is not valid")
	}
	return nil
}