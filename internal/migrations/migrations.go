package migrations

import (
	"fmt"
	"log"

	"jagajkn/internal/models"

	"gorm.io/gorm"
)

// Membuat enum types
func CreateEnumTypes(db *gorm.DB) error {
	// Jenis kelamin
	if err := db.Exec(`DO $$ BEGIN
		CREATE TYPE gender AS ENUM ('LAKI_LAKI', 'PEREMPUAN');
		EXCEPTION
		WHEN duplicate_object THEN null;
		END $$;`).Error; err != nil {
		return fmt.Errorf("error creating gender enum: %v", err)
	}

	// Kelas BPJS
	if err := db.Exec(`DO $$ BEGIN
		CREATE TYPE kelas_bpjs AS ENUM ('KELAS_1', 'KELAS_2', 'KELAS_3');
		EXCEPTION
		WHEN duplicate_object THEN null;
		END $$;`).Error; err != nil {
		return fmt.Errorf("error creating kelas_bpjs enum: %v", err)
	}

	// Jenis rawat
	if err := db.Exec(`DO $$ BEGIN
		CREATE TYPE jenis_rawat AS ENUM ('RAWAT_JALAN', 'RAWAT_INAP', 'RAWAT_DARURAT');
		EXCEPTION
		WHEN duplicate_object THEN null;
		END $$;`).Error; err != nil {
		return fmt.Errorf("error creating jenis_rawat enum: %v", err)
	}

	// Status pulang
	if err := db.Exec(`DO $$ BEGIN
		CREATE TYPE status_pulang AS ENUM ('SEMBUH', 'RUJUK', 'PULANG_PAKSA', 'MENINGGAL');
		EXCEPTION
		WHEN duplicate_object THEN null;
		END $$;`).Error; err != nil {
		return fmt.Errorf("error creating status_pulang enum: %v", err)
	}

	return nil
}


func RunMigrations(db *gorm.DB) error {
	log.Println("Starting database migrations...")


	if err := CreateEnumTypes(db); err != nil {
		return fmt.Errorf("error creating enum types: %v", err)
	}


	err := db.AutoMigrate(
		&models.User{},
		&models.RecordKesehatan{},
		&models.ResepObat{},
	)
	if err != nil {
		return fmt.Errorf("error running migrations: %v", err)
	}

	log.Println("Database migrations completed successfully")
	return nil
}