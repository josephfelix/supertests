package actions

import (
	"net/http"
	"supertests/models"

	"github.com/gobuffalo/buffalo"
)

func TestIndex(context buffalo.Context) error {
	test := models.Test{}
	slug := context.Param("slug")
	err := models.DB.Where("slug = ?", slug).First(&test)

	if err != nil {
		return context.Error(http.StatusNotFound, err)
	}

	context.Set("test", test)

	return context.Render(200, renderer.HTML("test/index.html"))
}

func TestResult(context buffalo.Context) error {
	return context.Render(http.StatusOK, renderer.HTML("test/test.html"))
}
