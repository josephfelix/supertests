package actions

import (
	"net/http"
	"strings"
	"supertests/lib"
	"supertests/models"
	"supertests/quiz"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
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

func TestLoading(context buffalo.Context) error {
	referer := context.Request().Referer()
	slug := context.Param("slug")

	if !strings.Contains(referer, slug) {
		return context.Redirect(http.StatusTemporaryRedirect,
			"tPath()", render.Data{"slug": slug})
	}

	context.Set("slug", slug)

	return context.Render(http.StatusOK, renderer.HTML("test/loading.html"))
}

func TestProcess(context buffalo.Context) error {
	slug := context.Param("slug")
	session := context.Session()
	// outputPath, _ := filepath.Abs("public/r")
	test := models.Test{}
	err := models.DB.Where("slug = ?", slug).First(&test)

	if err != nil {
		return context.Error(http.StatusNotFound, err)
	}

	plugin := lib.QuizLoad(test.Plugin)

	quizuser, err := plugin.Lookup("User")

	if err != nil {
		panic(err)
	}

	*quizuser.(*quiz.User) = quiz.User{
		Id:    session.Get("userid").(string),
		Name:  session.Get("name").(string),
		Email: session.Get("email").(string),
		Photo: session.Get("photo").(string),
	}

	render, err := plugin.Lookup("Render")

	if err != nil {
		panic(err)
	}

	image, _ := render.(func() (string, error))()

	return context.Render(http.StatusOK, renderer.JSON(image))
}
