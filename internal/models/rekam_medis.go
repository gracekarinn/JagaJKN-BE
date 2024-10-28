package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JenisRawat string
type StatusPulang string

const (
	RawatJalan   JenisRawat = "RAWAT_JALAN"
	RawatInap    JenisRawat = "RAWAT_INAP"
	RawatDarurat JenisRawat = "RAWAT_DARURAT"

	Sembuh      StatusPulang = "SEMBUH"
	Rujuk       StatusPulang = "RUJUK"
	PulangPaksa StatusPulang = "PULANG_PAKSA"
	Meninggal   StatusPulang = "MENINGGAL"
)

// Untuk membuat record kesehatan baru, kita memerlukan input dari pengguna
type RecordInput struct {
	NoSEP          string     `json:"noSEP" binding:"required"`
	JenisRawat     JenisRawat `json:"jenisRawat" binding:"required"`
	DiagnosaAwal   string     `json:"diagnosaAwal" binding:"required"`
	DiagnosaPrimer string     `json:"diagnosaPrimer" binding:"required"`
	IcdX           string     `json:"icdX" binding:"required"`
	Tindakan       string     `json:"tindakan" binding:"required"`
}

// Rekor kesehatan yang akan disimpan dalam blockchain
type RecordKesehatan struct {
    NoSEP            string        `gorm:"type:varchar(20);primary_key" json:"noSEP"`
    CreatedAt        time.Time     `json:"createdAt"`
    UpdatedAt        time.Time     `json:"updatedAt"`
    UserNIK          string        `gorm:"type:varchar(16);not null" json:"userNIK"`
    User             User          `gorm:"foreignKey:UserNIK" json:"user,omitempty"`
    TanggalMasuk     time.Time     `gorm:"not null" json:"tanggalMasuk"`
    TanggalKeluar    *time.Time    `json:"tanggalKeluar,omitempty"`
    JenisRawat       JenisRawat    `gorm:"type:jenis_rawat;not null" json:"jenisRawat"`
    DiagnosaAwal     string        `gorm:"type:text;not null" json:"diagnosaAwal"`
    DiagnosaPrimer   string        `gorm:"type:text;not null" json:"diagnosaPrimer"`
    DiagnosaSekunder *string       `gorm:"type:text" json:"diagnosaSekunder,omitempty"`
    IcdX             string        `gorm:"type:varchar(10);not null" json:"icdX"`
    Tindakan         string        `gorm:"type:text;not null" json:"tindakan"`
    StatusPulang     *StatusPulang `gorm:"type:status_pulang" json:"statusPulang,omitempty"`
    // Untuk menyimpan hash dari record sebelumnya
	BlockchainVerified bool `gorm:"-" json:"blockchainVerified"`
    HashPrevious   *string `gorm:"type:varchar(64)" json:"hashPrevious,omitempty"`
    HashCurrent    string  `gorm:"type:varchar(64);not null" json:"hashCurrent"`
    RetentionYears int     `gorm:"default:5" json:"retentionYears"`
}

type ResepObat struct {
	ID                string          `gorm:"type:varchar(36);primary_key" json:"id"`
	RecordKesehatanID string          `gorm:"type:varchar(36);not null" json:"recordKesehatanId"`
	RecordKesehatan   RecordKesehatan `gorm:"foreignKey:RecordKesehatanID" json:"recordKesehatan,omitempty"`
	KodeObat          string          `gorm:"type:varchar(20);not null" json:"kodeObat"`
	NamaObat          string          `gorm:"type:varchar(255);not null" json:"namaObat"`
	Dosis             string          `gorm:"type:varchar(100);not null" json:"dosis"`
	Jumlah            int             `gorm:"not null" json:"jumlah"`
	AturanPakai       string          `gorm:"type:text;not null" json:"aturanPakai"`
	HashData          string          `gorm:"type:varchar(64);not null" json:"hashData"`
}

func (r *RecordKesehatan) BeforeCreate(tx *gorm.DB) error {
	r.TanggalMasuk = time.Now()
	return nil
}

func (r *ResepObat) BeforeCreate(tx *gorm.DB) error {
	r.ID = uuid.New().String()
	return nil
}


func (r *RecordKesehatan) ToBlockchainRecord() map[string]interface{} {
	return map[string]interface{}{
		"noSEP":            r.NoSEP,
		"userNIK":          r.UserNIK,
		"tanggalMasuk":     r.TanggalMasuk,
		"tanggalKeluar":    r.TanggalKeluar,
		"jenisRawat":       r.JenisRawat,
		"diagnosaAwal":     r.DiagnosaAwal,
		"diagnosaPrimer":   r.DiagnosaPrimer,
		"diagnosaSekunder": r.DiagnosaSekunder,
		"icdX":             r.IcdX,
		"tindakan":         r.Tindakan,
		"statusPulang":     r.StatusPulang,
		"hashPrevious":     r.HashPrevious,
		"hashCurrent":      r.HashCurrent,
		"createdAt":        r.CreatedAt,
		"updatedAt":        r.UpdatedAt,
	}
}