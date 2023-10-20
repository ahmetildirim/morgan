package user

import (
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidPassword = errors.New("invalid password")
)

type User struct {
	ID             uuid.UUID
	Email          string
	HashedPassword []byte
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewUser(email string, password string) (*User, error) {
	if !isValidEmail(email) {
		return nil, ErrInvalidEmail
	}

	user := &User{
		ID:        uuid.New(),
		Email:     email,
		CreatedAt: time.Now(),
	}
	err := user.setPassword(password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) setPassword(password string) error {
	if len(password) < 8 {
		return ErrInvalidPassword
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.HashedPassword = hashedPassword
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword(u.HashedPassword, []byte(password))
	return err == nil
}

func isValidEmail(email string) bool {
	// This is a simple email validation regex, it may not cover all cases
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
