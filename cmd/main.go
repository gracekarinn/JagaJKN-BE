package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"

	"jagajkn/internal/handlers"
	"jagajkn/internal/middleware"
	"jagajkn/internal/repository"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	userRepo := repository.NewUserRepository(conn)
	userHandler := handlers.NewUserHandler(userRepo, os.Getenv("JWT_SECRET"))


	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("JagaJKN API is running!"))
    }).Methods("GET")


	r.HandleFunc("/api/auth/signup", userHandler.SignUp).Methods("POST")
	r.HandleFunc("/api/auth/login", userHandler.Login).Methods("POST")

	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.AuthMiddleware(os.Getenv("JWT_SECRET")))
	protected.HandleFunc("/user/profile", userHandler.GetProfile).Methods("GET")


	if err := createTables(conn); err != nil {
		log.Fatalf("Error creating tables: %v\n", err)
	}


	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func createTables(conn *pgx.Conn) error {
	userTable := `
		CREATE TABLE IF NOT EXISTS "User" (
			id TEXT PRIMARY KEY,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			"NIK" TEXT UNIQUE NOT NULL,
			nama_lengkap TEXT NOT NULL,
			tanggal_lahir TIMESTAMP NOT NULL,
			no_telp TEXT NOT NULL,
			email TEXT UNIQUE,
			password TEXT NOT NULL,
			no_kartu_jkn TEXT UNIQUE NOT NULL,
			jenis_kelamin TEXT NOT NULL DEFAULT 'LAKI_LAKI',
			alamat TEXT NOT NULL DEFAULT 'Alamat belum diisi', 
			faskes_tingkat1 TEXT NOT NULL DEFAULT 'Faskes belum diisi',
			kelas_perawatan TEXT NOT NULL DEFAULT 'KELAS_3'
		);
	`
	
	recordTable := `
		CREATE TABLE IF NOT EXISTS "RecordKesehatan" (
			id TEXT PRIMARY KEY,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			user_id TEXT NOT NULL,
			FOREIGN KEY (user_id) REFERENCES "User"(id)
		);
	`

	if _, err := conn.Exec(context.Background(), userTable); err != nil {
		return fmt.Errorf("error creating User table: %v", err)
	}

	if _, err := conn.Exec(context.Background(), recordTable); err != nil {
		return fmt.Errorf("error creating RecordKesehatan table: %v", err)
	}

	return nil
}