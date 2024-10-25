package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"jagajkn/internal/middleware"
	"jagajkn/internal/models"
	"jagajkn/internal/repository"
	"jagajkn/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	repo      *repository.UserRepository
	jwtSecret string
}

func NewUserHandler(repo *repository.UserRepository, jwtSecret string) *UserHandler {
	return &UserHandler{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req models.UserSignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	existingUser, err := h.repo.GetUserByNIK(r.Context(), req.NIK)
	if err != nil && err.Error() != "no rows in result set" {
		http.Error(w, "Error checking existing user", http.StatusInternalServerError)
		return
	}
	if existingUser != nil {
		http.Error(w, "User with this NIK already exists", http.StatusConflict)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error processing password", http.StatusInternalServerError)
		return
	}

	user := &models.User{
		ID:             utils.GenerateID(),
		CreatedAt:      time.Now(),
		NIK:            req.NIK,
		NamaLengkap:    req.NamaLengkap,
		NoTelp:         req.NoTelp,
		Email:          req.Email,
		Password:       string(hashedPassword),
		NoKartuJKN:     req.NoKartuJKN,
		JenisKelamin:   models.LAKI_LAKI, 
		TanggalLahir:   time.Now(),      
		Alamat:         "",              
		FaskesTingkat1: "",              
		KelasPerawatan: models.KELAS_3,   
	}


	if err := h.repo.CreateUser(r.Context(), user); err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}


	token, err := utils.GenerateToken(user, h.jwtSecret)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}


	response := models.UserResponse{
		User:  user,
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}


func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}


	user, err := h.repo.GetUserByNIK(r.Context(), req.NIK)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}


	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}


	token, err := utils.GenerateToken(user, h.jwtSecret)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}


	response := models.UserResponse{
		User:  user,
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}


func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {

	userID, err := middleware.GetUserIDFromContext(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}


	user, err := h.repo.GetUserByNIK(r.Context(), userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	response := models.UserResponse{
		User: user,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}