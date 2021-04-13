package actions

import (
	"net/http"
	"supertests/models"

	"github.com/gobuffalo/buffalo"
)

type LoginResult struct {
	Status bool `json:"status"`
}

func LoginHandler(context buffalo.Context) error {
	user := models.User{}

	if err := context.Bind(&user); err != nil {
		return err
	}

	exists, err := models.DB.Where("id = ?", user.ID).Exists(&user)

	if err != nil {
		return err
	}

	if !exists {
		err := models.DB.Create(&user)

		if err != nil {
			return err
		}
	}

	// TODO: Register user session

	return context.Render(http.StatusOK, renderer.JSON(LoginResult{Status: true}))
}

func LogoutHandler(context buffalo.Context) error {
	return context.Render(http.StatusOK, renderer.HTML("test/index.html"))
}
