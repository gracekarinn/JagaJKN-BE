package models

import (
	"time"
)

type Gender string

const (
	LAKI_LAKI Gender = "Laki-laki"
	PEREMPUAN Gender = "Perempuan"
)

type KelasBPJS string 

const (
	KELAS_1 KelasBPJS = "Kelas_1"
	KELAS_2 KelasBPJS = "Kelas_2"
	KELAS_3 KelasBPJS = "Kelas_3"
)

type User struct {
	ID              string     `json:"id"`
	CreatedAt       time.Time  `json:"createdAt"`
	NIK             string     `json:"nik"`
	NamaLengkap     string     `json:"namaLengkap"`
	TanggalLahir    time.Time  `json:"tanggalLahir"`
	NoTelp          string     `json:"noTelp"`
	Email           *string    `json:"email,omitempty"`
	Password        string     `json:"-"` 
	NoKartuJKN      string     `json:"noKartuJKN"`
	JenisKelamin    Gender     `json:"jenisKelamin"`
	Alamat          string     `json:"alamat"`
	FaskesTingkat1  string     `json:"faskesTingkat1"`
	KelasPerawatan  KelasBPJS  `json:"kelasPerawatan"`
}

type UserSignupRequest struct {
	NIK          string  `json:"nik" validate:"required,len=16"`
	NamaLengkap  string  `json:"namaLengkap" validate:"required"`
	NoTelp       string  `json:"noTelp" validate:"required"`
	Email        *string `json:"email,omitempty" validate:"omitempty,email"`
	Password     string  `json:"password" validate:"required,min=8"`
	NoKartuJKN   string  `json:"noKartuJKN" validate:"required"`
}

type UserLoginRequest struct {
	NIK      string `json:"nik" validate:"required,len=16"`
	Password string `json:"password" validate:"required"`
}

type UserUpdateRequest struct {
	NamaLengkap    string  `json:"namaLengkap" validate:"required"`
	PassportNumber string  `json:"passportNumber" validate:"required"`
}

type UserResponse struct {
	User *User `json:"user"`
	Token string `json:"token,omitempty"`
}
type RecordKesehatan struct {
	ID              string     `json:"id"`
	CreatedAt       time.Time  `json:"createdAt"`
	UserID          string     `json:"userID"`
	User 		    User       `json:"user"`
	// Nanti ditambahin
}

