package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type TingkatFaskes string

const (
    TingkatSatu  TingkatFaskes = "TINGKAT_1"
    TingkatDua   TingkatFaskes = "TINGKAT_2"
    TingkatTiga  TingkatFaskes = "TINGKAT_3"
)

type Faskes struct {
    KodeFaskes string       `gorm:"type:varchar(20);primary_key" json:"kodeFaskes"`
    Nama       string       `gorm:"type:varchar(255);not null" json:"nama"`
    Alamat     string       `gorm:"type:text;not null" json:"alamat"`
    NoTelp     string       `gorm:"type:varchar(15);not null" json:"noTelp"`
    Tingkat    TingkatFaskes `gorm:"type:tingkat_faskes;not null" json:"tingkat"`
    Email      string       `gorm:"type:varchar(255);unique;not null" json:"email"`
    Password   string       `gorm:"type:varchar(255);not null" json:"-"` 
    CreatedAt  time.Time    `json:"createdAt"`
    UpdatedAt  time.Time    `json:"updatedAt"`
}

type FaskesInput struct {
    KodeFaskes string       `json:"kodeFaskes" binding:"required"`
    Nama       string       `json:"nama" binding:"required"`
    Alamat     string       `json:"alamat" binding:"required"`
    NoTelp     string       `json:"noTelp" binding:"required"`
    Tingkat    TingkatFaskes `json:"tingkat" binding:"required"`
    Email      string       `json:"email" binding:"required,email"`
    Password   string       `json:"password" binding:"required,min=6"`
}

type FaskesLoginInput struct {
    KodeFaskes string `json:"kodeFaskes" binding:"required"`
    Password   string `json:"password" binding:"required"`
}


func (f *Faskes) HashPassword() error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(f.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    f.Password = string(hashedPassword)
    return nil
}


func (f *Faskes) CheckPassword(password string) error {
    return bcrypt.CompareHashAndPassword([]byte(f.Password), []byte(password))
}


func (f *Faskes) ToJSON() map[string]interface{} {
    return map[string]interface{}{
        "kodeFaskes": f.KodeFaskes,
        "nama":       f.Nama,
        "alamat":     f.Alamat,
        "noTelp":     f.NoTelp,
        "tingkat":    f.Tingkat,
        "email":      f.Email,
        "createdAt":  f.CreatedAt,
        "updatedAt":  f.UpdatedAt,
    }
}


func (f *Faskes) BeforeCreate(tx *gorm.DB) error {
    return f.HashPassword()
}