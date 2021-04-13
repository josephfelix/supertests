package middleware

import (
	"github.com/gobuffalo/buffalo"
)

func Login(next buffalo.Handler) buffalo.Handler {
	return func(context buffalo.Context) error {

		context.Set("isloggedin", false)
		context.Set("photo", "")
		context.Set("name", "Malakias")

		err := next(context)
		return err
	}
}
