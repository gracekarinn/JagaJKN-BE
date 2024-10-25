package repository

import (
	"context"
	"jagajkn/internal/models"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	conn *pgx.Conn
}

func NewUserRepository(conn *pgx.Conn) *UserRepository {
	return &UserRepository{
		conn: conn,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO "User" (
			id, created_at, "NIK", nama_lengkap, tanggal_lahir, 
			no_telp, email, password, no_kartu_jkn, jenis_kelamin,
			alamat, faskes_tingkat1, kelas_perawatan
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`

	_, err := r.conn.Exec(ctx, query,
		user.ID,
		user.CreatedAt,
		user.NIK,
		user.NamaLengkap,
		user.TanggalLahir,
		user.NoTelp,
		user.Email,
		user.Password,
		user.NoKartuJKN,
		user.JenisKelamin,
		user.Alamat,
		user.FaskesTingkat1,
		user.KelasPerawatan,
	)

	return err
}

func (r *UserRepository) GetUserByNIK(ctx context.Context, nik string) (*models.User, error) {
	query := `
		SELECT id, created_at, "NIK", nama_lengkap, tanggal_lahir,
			no_telp, email, password, no_kartu_jkn, jenis_kelamin,
			alamat, faskes_tingkat1, kelas_perawatan
		FROM "User"
		WHERE "NIK" = $1
	`

	var user models.User
	err := r.conn.QueryRow(ctx, query, nik).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.NIK,
		&user.NamaLengkap,
		&user.TanggalLahir,
		&user.NoTelp,
		&user.Email,
		&user.Password,
		&user.NoKartuJKN,
		&user.JenisKelamin,
		&user.Alamat,
		&user.FaskesTingkat1,
		&user.KelasPerawatan,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
		SELECT id, created_at, "NIK", nama_lengkap, tanggal_lahir,
			no_telp, email, password, no_kartu_jkn, jenis_kelamin,
			alamat, faskes_tingkat1, kelas_perawatan
		FROM "User"
		WHERE email = $1
	`

	var user models.User
	err := r.conn.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.NIK,
		&user.NamaLengkap,
		&user.TanggalLahir,
		&user.NoTelp,
		&user.Email,
		&user.Password,
		&user.NoKartuJKN,
		&user.JenisKelamin,
		&user.Alamat,
		&user.FaskesTingkat1,
		&user.KelasPerawatan,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}