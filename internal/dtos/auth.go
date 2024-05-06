package dtos

import (
	"github.com/iBoBoTi/standup-management-tool/internal/validator"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (req *LoginRequest) Validate(v *validator.Validator) bool {
	v.Check(req.Email != "", "email", "must be provided")
	v.Check(req.Password != "", "password", "must be provided")

	return v.Valid()
}

// CheckPassword checks if plainPassword matches hashedPassword
func (req *LoginRequest) CheckPassword(hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))
}

type LoginResponse struct {
	LoginData interface{}
}
