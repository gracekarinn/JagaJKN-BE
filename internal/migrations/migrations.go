package migrations

import (
	"fmt"
	"jagajkn/internal/models"
	"log"

	"gorm.io/gorm"
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

    return nil
}


func RunMigrations(db *gorm.DB) error {
    log.Println("Starting database migrations...")
    
    if err := CreateEnumTypes(db); err != nil {
        return fmt.Errorf("error creating enum types: %v", err)
    }

    err := db.AutoMigrate(
        &models.User{},
        &models.Admin{},
        &models.RecordKesehatan{},
        &models.ResepObat{},
        &models.Faskes{},
    )
    if err != nil {
        return fmt.Errorf("error running migrations: %v", err)
    }

    var admin models.Admin
    if db.Where("email = ?", "admin@jagajkn.com").First(&admin).Error != nil {
        log.Println("Creating default admin...")
        defaultAdmin := models.Admin{
            Email:    "admin@jagajkn.com",
            Password: "test123",
        }
        
        if err := defaultAdmin.HashPassword(); err != nil {
            log.Printf("Error hashing password: %v", err)
            return fmt.Errorf("error hashing admin password: %v", err)
        }
        
        if err := db.Create(&defaultAdmin).Error; err != nil {
            log.Printf("Error creating admin: %v", err)
            return fmt.Errorf("error creating default admin: %v", err)
        }
        log.Println("Default admin account created successfully")
    }

    log.Println("Database migrations completed successfully")
    return nil
}