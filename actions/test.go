package actions

import (
	"fmt"
	"image"
	"net/http"
	"strings"
	"supertests/lib"
	"supertests/models"
	"supertests/quiz"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	uuidv4 "github.com/satori/go.uuid"
)

type TestResponse struct {
	Status bool   `json:"status"`
	Hash   string `json:"hash"`
}

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
	test := models.Test{}
	imagebuilder := lib.ImageBuilder{}
	err := models.DB.Where("slug = ?", slug).First(&test)

	if err != nil {
		return context.Error(http.StatusNotFound, err)
	}

	plugin := lib.QuizLoad(test.Plugin)
	quizuser, err := plugin.Lookup("User")

	if err != nil {
		return context.Error(http.StatusNotFound, err)
	}

	*quizuser.(*quiz.User) = quiz.User{
		Id:    session.Get("userid").(string),
		Name:  session.Get("name").(string),
		Email: session.Get("email").(string),
		Photo: session.Get("photo").(string),
	}

	render, err := plugin.Lookup("Render")

	if err != nil {
		return context.Error(http.StatusNotFound, err)
	}

	img, err := render.(func() (image.Image, error))()

	if err != nil {
		return context.Error(http.StatusNotFound, err)
	}

	hash := uuidv4.NewV4().String()
	filename := hash + ".jpg"
	err = imagebuilder.Save(img, fmt.Sprintf("public/r/%v", filename))

	if err != nil {
		return context.Error(http.StatusNotFound, err)
	}

	result := models.Result{
		ID:        uuidv4.NewV4().String(),
		UserID:    session.Get("userid").(string),
		Guid:      slug,
		Result:    hash,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := models.DB.Create(&result); err != nil {
		return context.Error(http.StatusNotFound, err)
	}

	response := TestResponse{Status: true, Hash: hash}

	return context.Render(http.StatusOK, renderer.JSON(response))
}
