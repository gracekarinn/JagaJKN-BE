package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"jagajkn/internal/models"
	"jagajkn/internal/repository"
)

type UserService struct {
    repo      *repository.UserRepository
    jwtSecret string
}

func NewUserService(repo *repository.UserRepository, jwtSecret string) *UserService {
    return &UserService{
        repo:      repo,
        jwtSecret: jwtSecret,
    }
}

func (s *UserService) Register(input *models.UserSignupInput) (*models.User, error) {
    existing, _ := s.repo.FindByNIK(input.NIK)
    if existing != nil {
        return nil, errors.New("user already exists")
    }

    user := &models.User{
        NIK:         input.NIK,
        NamaLengkap: input.NamaLengkap,
        NoTelp:      input.NoTelp,
        Email:       input.Email,
        Password:    input.Password,
    }

    if err := s.repo.Create(user); err != nil {
        return nil, err
    }

    return user, nil
}

func (s *UserService) Login(input *models.UserLoginInput) (string, error) {
    user, err := s.repo.FindByNIK(input.NIK)
    if err != nil {
        return "", errors.New("invalid credentials")
    }

    if err := user.CheckPassword(input.Password); err != nil {
        return "", errors.New("invalid credentials")
    }

    claims := jwt.MapClaims{
        "user_nik": user.NIK,
        "role": string(user.Role),  
        "exp":      time.Now().Add(time.Hour * 24).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte(s.jwtSecret))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func (s *UserService) GetUserByNIK(nik string) (*models.User, error) {
    user, err := s.repo.FindByNIK(nik)
    if err != nil {
        return nil, errors.New("user not found")
    }
    return user, nil
}