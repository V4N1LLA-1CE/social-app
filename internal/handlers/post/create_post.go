package post

import (
	"net/http"

	"github.com/V4N1LLA-1CE/netio"
	"github.com/V4N1LLA-1CE/social-app/internal/store"
)

type CreatePostPayload struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	// read request body into input
	var input CreatePostPayload
	err := netio.Read(w, r, &input)
	if err != nil {
		netio.Error(w, "error", err.Error, http.StatusBadRequest)
		return
	}

	userId := 1
	ctx := r.Context()
	post := &store.Post{
		Title:   input.Title,
		Content: input.Content,
		UserID:  int64(userId),
		Tags:    input.Tags,
	}

	if err = h.store.Posts.Create(ctx, post); err != nil {
		netio.Error(w, "error", err.Error(), http.StatusInternalServerError)
		return
	}

	if err = netio.Write(w, http.StatusCreated, netio.Envelope{"post": post}, nil); err != nil {
		netio.Error(w, "error", err.Error(), http.StatusInternalServerError)
		return
	}
}
