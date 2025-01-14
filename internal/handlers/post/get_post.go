package post

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/V4N1LLA-1CE/netio"
	"github.com/V4N1LLA-1CE/social-app/internal/store"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) GetPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "postID"), 10, 64)
	if err != nil {
		netio.Error(w, "error", http.StatusInternalServerError, nil)
		return
	}

	ctx := r.Context()

	post, err := h.store.Posts.GetById(ctx, id)
	if err != nil {
		switch {

		case errors.Is(err, store.ErrNotFound):
			netio.Error(w, "error", http.StatusNotFound, nil)
			return

		default:
			netio.Error(w, "error", http.StatusInternalServerError, nil)
			return
		}
	}

	comments, err := h.store.Comments.GetByPostID(ctx, post.ID)
	if err != nil {
		netio.Error(w, "error", http.StatusInternalServerError, nil)
		return
	}

	post.Comments = comments

	// build response
	e := netio.Envelope{
		"post": post,
	}

	if err = netio.Write(w, http.StatusOK, e, nil); err != nil {
		netio.Error(w, "error", http.StatusInternalServerError, nil)
		return
	}
}
