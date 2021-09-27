package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"okki.hu/garric/ppnext/config"
	"okki.hu/garric/ppnext/model"
	"okki.hu/garric/ppnext/mw"
)

// ShowLogin sends the login page template to the client, and
// also binds optional query parameters to the template - this
// is useful for pre-filling certain form fields for share URLs
func ShowLogin(c *gin.Context) {

	var qp model.LoginQueryParams
	c.ShouldBindQuery(&qp)

	h := gin.H{
		"room":  qp.Room,
		"name":  qp.Name,
		"valid": qp.Valid,
		"email": config.Support,
	}

	// check if user is logged in
	user, ok := c.Get("user")
	if ok {
		h["name"] = user
		h["state"] = "readonly"
	}

	c.HTML(http.StatusOK, "login.html", h)
}

// HandleLogin checks if the user
func HandleLogin(c *gin.Context) {

	var form model.LoginForm
	c.ShouldBind(&form)

	// check if user is logged in
	user, ok := c.Get("user")
	if !ok {
		// ensure user name is not taken
		exists, err := config.Repository.Exists(form.Name)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		if exists {
			loc := fmt.Sprintf("/login?room=%s&name=%s&valid=invalid", form.Room, form.Name)
			c.Redirect(http.StatusFound, loc)
			return
		}

		// set cookie with user name
		mw.SetAuthCookie(c, form.Name)
		user = form.Name
	}

	name := user.(string)

	// if needed, add user to the room
	room, err := config.Repository.Load(form.Room)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	if _, ok := room.Votes[name]; !ok {
		room.RegisterVote(&model.Vote{
			User: name,
			Vote: config.Nothing,
		})
	}
	config.Repository.Save(room)

	loc := fmt.Sprintf("/rooms/%s", form.Room)
	c.Redirect(http.StatusFound, loc)
}
