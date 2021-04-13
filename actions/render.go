package actions

import (
	"supertests/helpers"

	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr/v2"
)

var renderer *render.Engine
var assetsBox = packr.New("app:assets", "../public")

func init() {
	renderer = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.plush.html",

		// Box containing all of the templates:
		TemplatesBox: packr.New("app:templates", "../templates"),
		AssetsBox:    assetsBox,

		// Add template helpers here:
		Helpers: render.Helpers{
			// for non-bootstrap form helpers uncomment the lines
			// below and import "github.com/gobuffalo/helpers/forms"
			"Session": helpers.Session,
			// forms.FormForKey:  forms.FormFor,
		},
	})
}
