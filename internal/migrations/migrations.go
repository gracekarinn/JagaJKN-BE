package migrations

import (
	"fmt"
	"jagajkn/internal/models"
	"log"
	"os"

	"gorm.io/gorm"
)

var (
    infoLog  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
    errorLog = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)
)

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

    // Add tingkat_faskes enum
    if err := db.Exec(`DO $$ BEGIN
        CREATE TYPE tingkat_faskes AS ENUM ('TINGKAT_1', 'TINGKAT_2', 'TINGKAT_3');
        EXCEPTION
        WHEN duplicate_object THEN null;
        END $$;`).Error; err != nil {
        return fmt.Errorf("error creating tingkat_faskes enum: %v", err)
    }

    if err := db.Exec(`DO $$ BEGIN
        CREATE TYPE status_transfer AS ENUM ('PENDING', 'ACCEPTED', 'REJECTED');
        EXCEPTION
        WHEN duplicate_object THEN null;
        END $$;`).Error; err != nil {
        return fmt.Errorf("error creating status_transfer enum: %v", err)
    }

    return nil
}


func RunMigrations(db *gorm.DB) error {
    // db.Exec("DROP TABLE IF EXISTS rekam_medis_transfers CASCADE")
    // db.Exec("DROP TABLE IF EXISTS resep_obat CASCADE")
    // db.Exec("DROP TABLE IF EXISTS record_kesehatans CASCADE")
    // db.Exec("DROP TABLE IF EXISTS faskes CASCADE")
    // db.Exec("DROP TABLE IF EXISTS users CASCADE")
    // db.Exec("DROP TABLE IF EXISTS admins CASCADE")

    // db.Exec("DROP TYPE IF EXISTS status_transfer CASCADE")
    // db.Exec("DROP TYPE IF EXISTS tingkat_faskes CASCADE")
    // db.Exec("DROP TYPE IF EXISTS jenis_rawat CASCADE")
    // db.Exec("DROP TYPE IF EXISTS status_pulang CASCADE")
    // db.Exec("DROP TYPE IF EXISTS gender CASCADE")
    // db.Exec("DROP TYPE IF EXISTS kelas_bpjs CASCADE")

    for _, model := range []interface{}{
        &models.Admin{},
        &models.User{},
        &models.Faskes{},
        &models.RecordKesehatan{},
        &models.ResepObat{},
        &models.RekamMedisTransfer{},
    } {
        if err := db.AutoMigrate(model); err != nil {
            return fmt.Errorf("error migrating %T: %v", model, err)
        }
        infoLog.Printf("✅ Successfully migrated %T", model)
    }

    var admin models.Admin
    if db.Where("email = ?", "admin@jagajkn.com").First(&admin).Error != nil {
        infoLog.Println("Creating default admin account...")
        defaultAdmin := models.Admin{
            Email:    "admin@jagajkn.com",
            Password: "test123",
        }
        if err := defaultAdmin.HashPassword(); err != nil {
            return fmt.Errorf("error hashing admin password: %v", err)
        }
        if err := db.Create(&defaultAdmin).Error; err != nil {
            return fmt.Errorf("error creating default admin: %v", err)
        }
        infoLog.Println("✅ Default admin account created successfully")
    }

    return nil
}