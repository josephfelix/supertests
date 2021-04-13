package middleware

import (
	"github.com/gobuffalo/buffalo"
)

func Login(next buffalo.Handler) buffalo.Handler {
	return func(context buffalo.Context) error {
		session := context.Session()

		context.Set("isloggedin", session.Get("loggedin"))
		context.Set("userid", session.Get("userid"))
		context.Set("photo", session.Get("photo"))
		context.Set("name", session.Get("name"))

		err := next(context)
		return err
	}
}
