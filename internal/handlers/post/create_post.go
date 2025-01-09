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
		netio.Error(w, "error", http.StatusBadRequest, nil)
		return
	}

	if input.Tags == nil {
		input.Tags = []string{}
	}

	// perform validation of payload
	v := netio.NewValidator()
	v.Check(len(input.Title) <= 100, "title", "must not exceed 100 characters")
	v.Check(len(input.Title) > 0, "title", "must not be empty")
	v.Check(len(input.Content) <= 1000, "content", "must not exceed 1000 characters")
	v.Check(len(input.Content) > 0, "content", "must not be empty")
	v.Check(!netio.HasDuplicates(input.Tags), "tags", "must not have duplicates")
	if !v.Valid() {
		netio.Error(w, "error", http.StatusUnprocessableEntity, v)
		return
	}

	ctx := r.Context()
	post := &store.Post{
		Title:   input.Title,
		Content: input.Content,
		UserID:  int64(1),
		Tags:    input.Tags,
	}

	if err = h.store.Posts.Create(ctx, post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		netio.Error(w, "error", http.StatusInternalServerError, nil)
		return
	}

	if err = netio.Write(w, http.StatusCreated, netio.Envelope{"post": post}, nil); err != nil {
		netio.Error(w, "error", http.StatusInternalServerError, nil)
		return
	}
}
