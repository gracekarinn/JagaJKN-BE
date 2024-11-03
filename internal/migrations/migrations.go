package migrations

import (
	"fmt"
	"jagajkn/internal/models"
	"log"
	"os"

	"gorm.io/gorm"
)

var (
    infoLog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
)

// EnumDefinition represents a PostgreSQL enum type
type EnumDefinition struct {
    name    string
    values  []string
}

var enums = []EnumDefinition{
    {
        name: "gender",
        values: []string{"LAKI_LAKI", "PEREMPUAN"},
    },
    {
        name: "kelas_bpjs",
        values: []string{"KELAS_1", "KELAS_2", "KELAS_3"},
    },
    {
        name: "jenis_rawat",
        values: []string{"RAWAT_JALAN", "RAWAT_INAP", "RAWAT_DARURAT"},
    },
    {
        name: "status_pulang",
        values: []string{"SEMBUH", "RUJUK", "PULANG_PAKSA", "MENINGGAL"},
    },
    {
        name: "tingkat_faskes",
        values: []string{"TINGKAT_1", "TINGKAT_2", "TINGKAT_3"},
    },
    {
        name: "status_transfer",
        values: []string{"PENDING", "ACCEPTED", "REJECTED"},
    },
}

func CreateEnumTypes(db *gorm.DB) error {
    for _, enum := range enums {
        query := fmt.Sprintf(`DO $$ 
        BEGIN
            IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = '%s') THEN
                CREATE TYPE %s AS ENUM ('%s');
            END IF;
        END $$;`, enum.name, enum.name, "', '"+join(enum.values, "', '"))

        if err := db.Exec(query).Error; err != nil {
            return fmt.Errorf("error creating %s enum: %v", enum.name, err)
        }
        infoLog.Printf("✅ Created/Verified enum: %s", enum.name)
    }
    return nil
}

// func dropExistingObjects(db *gorm.DB) error {
//     // Tables to drop
//     tables := []string{
//         "rekam_medis_transfers",
//         "resep_obat",
//         "record_kesehatans",
//         "faskes",
//         "users",
//         "admins",
//     }

//     // Drop tables
//     for _, table := range tables {
//         if err := db.Exec("DROP TABLE IF EXISTS " + table + " CASCADE").Error; err != nil {
//             return fmt.Errorf("error dropping table %s: %v", table, err)
//         }
//         infoLog.Printf("Dropped table: %s", table)
//     }

//     // Drop enum types
//     for _, enum := range enums {
//         if err := db.Exec("DROP TYPE IF EXISTS " + enum.name + " CASCADE").Error; err != nil {
//             return fmt.Errorf("error dropping enum %s: %v", enum.name, err)
//         }
//         infoLog.Printf("Dropped enum: %s", enum.name)
//     }

//     return nil
// }

func createDefaultAdmin(db *gorm.DB) error {
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
    } else {
        infoLog.Println("ℹ️ Default admin account already exists")
    }
    return nil
}

func RunMigrations(db *gorm.DB) error {
    infoLog.Println("Starting database migrations...")

    // // Drop existing objects
    // if err := dropExistingObjects(db); err != nil {
    //     return fmt.Errorf("error dropping existing objects: %v", err)
    // }

    // Create enum types
    if err := CreateEnumTypes(db); err != nil {
        return fmt.Errorf("error creating enum types: %v", err)
    }

    // Models to migrate
    models := []interface{}{
        &models.Admin{},
        &models.User{},
        &models.Faskes{},
        &models.RecordKesehatan{},
        &models.ResepObat{},
        &models.RekamMedisTransfer{},
    }

    // Run migrations
    for _, model := range models {
        if err := db.AutoMigrate(model); err != nil {
            return fmt.Errorf("error migrating %T: %v", model, err)
        }
        infoLog.Printf("✅ Successfully migrated %T", model)
    }

    // Create default admin account
    if err := createDefaultAdmin(db); err != nil {
        return fmt.Errorf("error creating default admin: %v", err)
    }

    infoLog.Println("✅ All migrations completed successfully!")
    return nil
}

// Helper function to join strings with a separator
func join(strs []string, sep string) string {
    if len(strs) == 0 {
        return ""
    }
    result := strs[0]
    for _, str := range strs[1:] {
        result += sep + str
    }
    return result
}