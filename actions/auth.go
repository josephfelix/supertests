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
	session := context.Session()

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

	session.Set("loggedin", true)
	session.Set("userid", user.ID)
	session.Set("email", user.Email)
	session.Set("photo", user.Photo)
	session.Set("name", user.Name)

	return context.Render(http.StatusOK, renderer.JSON(LoginResult{Status: true}))
}

func LogoutHandler(context buffalo.Context) error {
	context.Session().Clear()

	return context.Redirect(http.StatusTemporaryRedirect, "rootPath()")
}
