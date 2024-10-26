package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Gender string
type KelasBPJS string

const (
	LakiLaki  Gender = "LAKI_LAKI"
	Perempuan Gender = "PEREMPUAN"

	Kelas1 KelasBPJS = "KELAS_1"
	Kelas2 KelasBPJS = "KELAS_2"
	Kelas3 KelasBPJS = "KELAS_3"
)

type UserSignupInput struct {
	NIK         string  `json:"nik" binding:"required,len=16"`
	NamaLengkap string  `json:"namaLengkap" binding:"required"`
	NoTelp      string  `json:"noTelp" binding:"required"`
	Email       *string `json:"email"`
	Password    string  `json:"password" binding:"required,min=6"`
}

type UserLoginInput struct {
	NIK      string `json:"nik" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	NIK 	   string     `gorm:"type:varchar(16);primaryKey" json:"nik"`
	CreatedAt   time.Time  `json:"createdAt"`
	NamaLengkap string     `gorm:"type:varchar(255);not null" json:"namaLengkap"`
	NoTelp      string     `gorm:"type:varchar(20);not null" json:"noTelp"`
	Email       *string    `gorm:"type:varchar(255);uniqueIndex" json:"email,omitempty"`
	Password    string     `gorm:"type:varchar(255);not null" json:"-"`
	NoKartuJKN     *string    `gorm:"type:varchar(13);uniqueIndex" json:"noKartuJKN,omitempty"`
	JenisKelamin   *Gender    `gorm:"type:gender" json:"jenisKelamin,omitempty"`
	TanggalLahir   *time.Time `gorm:"type:date" json:"tanggalLahir,omitempty"`
	Alamat         *string    `gorm:"type:text" json:"alamat,omitempty"`
	FaskesTingkat1 *string    `gorm:"type:varchar(255)" json:"faskesTingkat1,omitempty"`
	KelasPerawatan *KelasBPJS `gorm:"type:kelas_bpjs" json:"kelasPerawatan,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    u.Password = string(hashedPassword)
    return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u *User) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"nik":            u.NIK,
		"namaLengkap":    u.NamaLengkap,
		"noTelp":         u.NoTelp,
		"email":          u.Email,
		"noKartuJKN":     u.NoKartuJKN,
		"jenisKelamin":   u.JenisKelamin,
		"tanggalLahir":   u.TanggalLahir,
		"alamat":         u.Alamat,
		"faskesTingkat1": u.FaskesTingkat1,
		"kelasPerawatan": u.KelasPerawatan,
	}
}