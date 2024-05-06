package dtos

import (
	"fmt"
	"regexp"
	"time"

	"github.com/google/uuid"
	"github.com/iBoBoTi/standup-management-tool/internal/validator"
	"golang.org/x/crypto/bcrypt"
)

const (
	MinPasswordLength = 8
	MaxPasswordLength = 255
)

const (
	AdminRole    string = "admin"
	EmployeeRole string = "employee"
)

var (
	EmailRgx = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

type Employee struct {
	ID              uuid.UUID `json:"id"`
	FirstName       string    `json:"first_name" binding:"required"`
	LastName        string    `json:"last_name" binding:"required"`
	Email           string    `json:"email" binding:"required"`
	Password        string    `json:"password,omitempty" binding:"required"`
	ConfirmPassword string    `json:"confirm_password,omitempty" binding:"required"`
	Company         string    `json:"company"`
	Role            string    `json:"role"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (u *Employee) Validate(v *validator.Validator) bool {
	v.Check(u.Email != "", "email", "email must not be blank")
	v.Check(IsEmail(u.Email), "email", "must be a valid email address")
	v.Check(len(u.Email) <= 200, "email", "must not be more than 200 bytes long")

	v.Check(u.Password != "", "password", "must not be blank")
	v.Check(len(u.Password) >= MinPasswordLength, "password", "must be at least 8 characters long")
	v.Check(len(u.Password) <= MaxPasswordLength, "password", "the password is too long")
	v.Check(u.Password == u.ConfirmPassword, "password", "password must be the same as confirm password")

	v.Check(u.FirstName != "", "first_name", "must not be blank")
	v.Check(len(u.FirstName) <= 255, "first_name", "must not be more than 50 bytes long")

	v.Check(u.LastName != "", "last_name", "must not be blank")
	v.Check(len(u.LastName) <= 255, "last_name", "must not be more than 50 bytes long")
	return v.Valid()
}

func (e *Employee) HashPassword() error {
	hashedPassword, err := hashPassword(e.Password)
	if err != nil {
		return err
	}
	e.Password = string(hashedPassword)
	return nil
}

type AdminCreateEmployeeDto struct {
	ID              uuid.UUID `json:"id"`
	FirstName       string    `json:"first_name" binding:"required"`
	LastName        string    `json:"last_name" binding:"required"`
	Email           string    `json:"email" binding:"required"`
	Company         string    `json:"company"`
	Role            string    `json:"role"`
	Password string `json:"password"`
	PasswordHash string `json:"-"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (a *AdminCreateEmployeeDto) HashPassword() error {
	hashedPassword, err := hashPassword(a.Password)
	if err != nil {
		return err
	}
	a.PasswordHash = string(hashedPassword)
	return nil
}

func (u *AdminCreateEmployeeDto) Validate(v *validator.Validator) bool {
	v.Check(u.Email != "", "email", "email must not be blank")
	v.Check(IsEmail(u.Email), "email", "must be a valid email address")
	v.Check(len(u.Email) <= 200, "email", "must not be more than 200 bytes long")

	v.Check(u.FirstName != "", "first_name", "must not be blank")
	v.Check(len(u.FirstName) <= 255, "first_name", "must not be more than 50 bytes long")

	v.Check(u.LastName != "", "last_name", "must not be blank")
	v.Check(len(u.LastName) <= 255, "last_name", "must not be more than 50 bytes long")
	return v.Valid()
}

// IsEmail returns true if a string is a valid email address.
func IsEmail(value string) bool {
	if len(value) > 254 {
		return false
	}

	return EmailRgx.MatchString(value)
}

func hashPassword(password string) (string ,error){
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password %v", err)
	}
	return string(hashedPassword), nil
}