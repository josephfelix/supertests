package actions

import (
	"net/http"
	"strconv"
	"supertests/models"

	"github.com/gobuffalo/buffalo"
)

func HomeHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, renderer.HTML("index/index.html"))
}

func HomeLoad(context buffalo.Context) error {
	limitParam := context.Param("limit")

	limit, err := strconv.Atoi(limitParam)

	if err != nil {
		return err
	}

	tests := &models.Tests{}
	query := models.DB.Order("created_at desc").Limit(limit)

	if err := query.All(tests); err != nil {
		return err
	}

	return context.Render(http.StatusOK, renderer.JSON(tests))
}
