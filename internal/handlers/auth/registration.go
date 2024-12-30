package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/V4N1LLA-1CE/netio"
	"github.com/V4N1LLA-1CE/social-app/internal/store"
	"github.com/V4N1LLA-1CE/social-app/internal/validators"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username          string `json:"username"`
		Email             string `json:"email"`
		PasswordPlaintext string `json:"password"`
	}

	err := netio.Read(w, r, &input)
	if err != nil {
		netio.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	// validate input
	v := netio.NewValidator()

	validators.ValidateUsername(v, input.Username)
	validators.ValidateEmail(v, input.Email)
	validators.ValidatePassword(v, input.PasswordPlaintext)

	if !v.Valid() {
		netio.Error(w, v.Errors, http.StatusUnprocessableEntity)
		return
	}

	// 5 second timeout on user creation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// hash password with bcrypt
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.PasswordPlaintext), 12)
	if err != nil {
		netio.Error(w, v.Errors, http.StatusInternalServerError)
		return
	}

	// create user
	user := &store.User{
		Username: input.Username,
		Email:    input.Email,
		Password: passwordHash,
	}

	if err := h.store.Users.Create(ctx, user); err != nil {
		netio.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
