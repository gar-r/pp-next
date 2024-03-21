package controller

import (
	"fmt"
	"net/http"

	"github.com/gar-r/ppnext/config"
	"github.com/gar-r/ppnext/model"
	"github.com/gar-r/ppnext/viewmodel"
	"github.com/gin-gonic/gin"
)

// ShowLogin sends the login page template to the client, and
// also binds optional query parameters to the template - this
// is useful for pre-filling certain form fields for share URLs
func ShowLogin(c *gin.Context) {

	var qp viewmodel.LoginQueryParams
	err := c.ShouldBindQuery(&qp)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	rooms, err := config.Repository.RoomCount()
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	users, err := config.Repository.UserCount()
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	h := gin.H{
		"room":  qp.Room,
		"name":  qp.Name,
		"valid": qp.Valid,
		"email": config.Support,
		"rooms": rooms,
		"users": users,
	}

	// check if user is logged in
	user, ok := c.Get("user")
	if ok {
		h["name"] = user
		h["state"] = "readonly"
	}

	c.HTML(http.StatusOK, "login.html.tmpl", h)
}

// HandleLogin logs in the user
func HandleLogin(c *gin.Context) {

	var form viewmodel.LoginForm
	err := c.ShouldBind(&form)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// check if user is logged in
	user, ok := c.Get("user")
	if !ok {
		// ensure username is not taken
		exists, err2 := config.Repository.Exists(form.Name)
		if err2 != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err2)
			return
		}
		if exists {
			loc := fmt.Sprintf("/login?room=%s&name=%s&valid=invalid", form.Room, form.Name)
			c.Redirect(http.StatusFound, loc)
			return
		}

		// set cookie with username
		SetAuthCookie(c, form.Name)
		user = form.Name
	}

	name := user.(string)

	// if needed, add user to the room
	room, err := config.Repository.Load(form.Room)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if _, ex := room.Votes[name]; !ex {
		room.RegisterVote(&model.Vote{
			User: name,
			Vote: model.Nothing,
		})
		err2 := config.Repository.Save(room)
		if err2 != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err2)
			return
		}
	}

	loc := fmt.Sprintf("/rooms/%s", form.Room)
	c.Redirect(http.StatusFound, loc)
}

func HandleLogout(c *gin.Context) {
	user, ok := c.Get("user")
	if ok {
		name := user.(string)
		err := config.Repository.Remove(name)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ClearAuthCookie(c)
	}
	c.Redirect(http.StatusFound, "/")
}
