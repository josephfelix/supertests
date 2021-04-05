package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
)

// TestHandler default implementation.
func TestIndex(context buffalo.Context) error {
	return context.Render(200, r.String(context.Param("uuid")))
}

func TestResult(context buffalo.Context) error {
	return context.Render(http.StatusOK, render.HTML("test/test.html"))
}
