package health

import (
	"net/http"
	"time"

	"github.com/V4N1LLA-1CE/netio"
)

type Handler struct {
	env string
}

func New(env string) *Handler {
	return &Handler{
		env: env,
	}
}

func (h *Handler) Check(w http.ResponseWriter, r *http.Request) {
	health := map[string]any{
		"status":      "alive",
		"environment": h.env,
		"timestamp":   time.Now(),
	}

	netio.Write(w, http.StatusOK, netio.Envelope{"health": health}, nil)
}
