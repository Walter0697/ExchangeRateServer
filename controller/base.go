package controller

import (
	"net/http"

	"github.com/go-chi/render"
)

type BaseRespond struct {
	Elapsed int64 `json:"-"`
}

func (br *BaseRespond) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type ErrResponse struct {
	BaseRespond
	ErrorText string `json:"error,omitempty"` // application-level error message, for debugging
}

func JSON(w http.ResponseWriter, r *http.Request, status int, v render.Renderer) error {
	render.Status(r, status)
	return render.Render(w, r, v)
}

func ErrorResp(w http.ResponseWriter, r *http.Request, status int, err error) error {
	render.Status(r, status)
	var errResp ErrResponse
	errResp.ErrorText = err.Error()
	return render.Render(w, r, &errResp)
}
