package controller

import (
	"fmt"
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

// Prot returns a middleware function, that protects pages which require
// authentication. Unauthenticated users get redirected to the login page.
func Prot() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, ok := c.Get("user")
		if !ok {
			room := c.Param("room")
			loc := fmt.Sprintf("/login?room=%s", room)
			c.Redirect(http.StatusFound, loc)
			c.Abort()
		}
	}
}

// Api returns a middleware function, that protects api routes which require
// authentication. Unauthenticated calls result in http.StatusUnauthorized.
func Api() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, ok := c.Get("user")
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

func SetAuthCookie(c *gin.Context, user string) {
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie(CookieName, user, CookieExpiry, "", consts.Domain, false, true)
}

func ClearAuthCookie(c *gin.Context) {
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie(CookieName, "", -1, "", consts.Domain, false, true)
}
