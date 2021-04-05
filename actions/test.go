package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// TestHandler default implementation.
func TestIndex(context buffalo.Context) error {
	return context.Render(200, renderer.String(context.Param("uuid")))
}

func TestResult(context buffalo.Context) error {
	return context.Render(http.StatusOK, renderer.HTML("test/test.html"))
}
