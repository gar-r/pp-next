package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"okki.hu/garric/ppnext/consts"
)

const CookieName = "ppnext-user"
const CookieExpiry = 60 * 60 * 6

// Auth returns a basic (unsecure) cookie based authentication
// middleware function.
// We look for the user cookie in the request, and if present,
// we set the 'user' context variable to its value.
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie(CookieName)
		if err == nil {
			c.Set("user", cookie)

		}
		c.Next()
	}
}

// Active returns a middleware function that can renew the user cookie.
// Normally the user auth cookie expires, but some user actions
// can extend this period. Any route that extends the user's
// activity window should use this middleware.
func Active() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, ok := c.Get("user")
		if ok {
			SetAuthCookie(c, user.(string))
		}
	}
}

// Prot returns a middleware function, that can protect routes
// that require authentication.
// Unauthenticated users get an HTTP/401 response.
func Prot() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, ok := c.Get("user")
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Next()
	}
}

func SetAuthCookie(c *gin.Context, user string) {
	c.SetCookie(CookieName, user, CookieExpiry, "", consts.Domain, false, true)
}
