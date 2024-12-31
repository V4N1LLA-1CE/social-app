package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/V4N1LLA-1CE/netio"
	"github.com/V4N1LLA-1CE/social-app/internal/store"
	"github.com/V4N1LLA-1CE/social-app/internal/validators"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email             string `json:"email"`
		PasswordPlaintext string `json:"password"`
	}

	// read json body into input struct
	err := netio.Read(w, r, &input)
	if err != nil {
		netio.Error(w, "error", err.Error(), http.StatusUnprocessableEntity)
		return
	}

	// validate input
	v := netio.NewValidator()

	validators.ValidateEmail(v, input.Email)
	validators.ValidatePassword(v, input.PasswordPlaintext)

	if !v.Valid() {
		netio.Error(w, "error", v.Errors, http.StatusUnprocessableEntity)
		return
	}

	// create user with default values
	// set email and attempt to set password
	user := &store.User{
		Email: input.Email,
	}
	if err := user.Password.Set(input.PasswordPlaintext); err != nil {
		netio.Error(w, "error", err.Error(), http.StatusInternalServerError)
	}

	// create user in db with 5 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := h.store.Users.Create(ctx, user); err != nil {
		netio.Error(w, "error", err.Error(), http.StatusInternalServerError)
		return
	}

	// send response
	if err := netio.Write(w, http.StatusCreated, netio.Envelope{"user": user}, nil); err != nil {
		netio.Error(w, "error", err.Error(), http.StatusInternalServerError)
		return
	}
}
