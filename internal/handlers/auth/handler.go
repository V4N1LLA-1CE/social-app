package auth

import "github.com/V4N1LLA-1CE/social-app/internal/store"

type Handler struct {
	store store.Store
}

func New(store store.Store) *Handler {
	return &Handler{
		store: store,
	}
}
