package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
    ID        uint      `gorm:"primary_key" json:"id"`
    Email     string    `gorm:"type:varchar(255);unique;not null" json:"email"`
    Password  string    `gorm:"type:varchar(255);not null" json:"-"` 
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}

type AdminLoginInput struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

func (a *Admin) HashPassword() error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    a.Password = string(hashedPassword)
    return nil
}

func (a *Admin) CheckPassword(password string) error {
    return bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password))
}

func (a *Admin) ToJSON() map[string]interface{} {
    return map[string]interface{}{
        "id":        a.ID,
        "email":     a.Email,
        "createdAt": a.CreatedAt,
        "updatedAt": a.UpdatedAt,
    }
}