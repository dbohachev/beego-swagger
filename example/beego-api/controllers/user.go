package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/dbohachev/beego-swagger/example/beego-api/models"
	"net/http"
	"strconv"
)

// Operations above user models
type UsersController struct {
	beego.Controller
}

// Get godoc
// @Summary Get list of all users
// @Description Get all users. No paging applied. Protected with basic auth (but on server side real protection does
// @Description not happen, logging only).
// @Tags users
// @Produce  json
// @Router /users [get]
// @Success 200 {array} models.User
func (u *UsersController) Get() {
	u.Data["json"] = models.GetAll()
	u.ServeJSON()
}

// GetByID godoc
// @Summary Get single user by it's ID
// @Description Get single user. ID of user should be specified. Error if no user found. Protected with basic auth
// @Description (but on server side real protection does not happen, logging only).
// @Tags users
// @Produce  json
// @Param  id path int true "User ID"
// @Router /users/{id} [get]
// @Success 200 {object} models.User
// @Failure 400 {object} models.RequestError
// @Failure 404 {object} models.RequestError
func (u *UsersController) GetByID() {
	strId := u.Ctx.Input.Param(":id")
	userId, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		u.serveBadRequest()
	} else {
		user, err := models.GetOne(userId)
		if err != nil {
			u.serveNotFound()
		} else {
			u.Data["json"] = user
			u.ServeJSON()
		}
	}
}

// Post godoc
// @Summary Add new user
// @Description Creates new user in our user set. User ID, if provided as part of user spec, is ignored.
// @Tags users
// @Security BasicAuth
// @Accept  json
// @Produce  json
// @Param user body models.User true "The user to add"
// @Success 200 {object} models.User
// @Failure 400 {object} models.RequestError
// @Router /users [post]
func (u *UsersController) Post() {
	user := models.User{}
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if err != nil {
		logs.Error("Could not parse 'Post' request payload, %s.", err)
		u.serveBadRequest()
	} else {
		user.ID = models.AddOne(user)
		u.Data["json"] = user
		u.ServeJSON()
	}
}

// Put godoc
// @Summary Update user
// @Description Update user with a json
// @Tags users
// @Security BasicAuth
// @Accept  json
// @Produce  json
// @Param  user body models.User true "The user to update"
// @Success 200 {object} models.User
// @Failure 400 {object} models.RequestError
// @Failure 404 {object} models.RequestError
// @Router /users [put]
func (u *UsersController) Put() {
	user := models.User{}
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if err != nil {
		logs.Error("Could not parse 'Put' request payload, %s.", err)
		u.serveBadRequest()
	} else {
		err = models.Update(user)
		if err != nil {
			logs.Error("Could not find user in 'Put' request, %s.", err)
			u.serveNotFound()
		} else {
			u.Data["json"] = user
			u.ServeJSON()
		}
	}
}

// Delete godoc
// @Summary Update user
// @Description Delete by user ID
// @Tags users
// @Security BasicAuth
// @Produce  json
// @Param  id path int true "User ID" Format(int64)
// @Success 200 {object} models.User
// @Failure 400 {object} models.RequestError
// @Router /users/{id} [delete]
func (u *UsersController) Delete() {
	strId := u.Ctx.Input.Param(":id")
	userId, err := strconv.ParseInt(strId, 10, 64)

	if err != nil {
		logs.Error("Could not parse 'Delete' request payload %s.", err)
		u.serveBadRequest()
	} else {
		// For simplicity, we don't block shared objects; simply checking if they exist.
		// In real life application the below code will be a race condition for globally shared
		// objects.
		user, err := models.Delete(userId)
		if err != nil {
			logs.Error("Could not delete user: %s.", err)
			u.serveNotFound()
		} else {
			u.Data["json"] = user
			u.ServeJSON()
		}
	}
}

// Simple wrapper around "request failed due to 400" functionality.
func (u *UsersController) serveBadRequest() {
	u.
		serveResponse("Malformed request data provided, could not continue.",
			http.StatusBadRequest)
}

// Simple wrapper around "request failed not found 404" functionality.
func (u *UsersController) serveNotFound() {
	u.
		serveResponse("Data specified with URL not found.",
			http.StatusNotFound)
}

// Simple wrapper around "request failed due to 500" functionality.
func (u *UsersController) serveInternalServerError() {
	u.
		serveResponse("Server encountered unexpected error while processing your request.",
			http.StatusInternalServerError)
}

// Carry out real response write, with specified data and response string. For simplicity, only string
// and used only from error handlers.
func (u *UsersController) serveResponse(message string, status int) {
	u.Ctx.ResponseWriter.WriteHeader(status)
	u.Data["json"] = models.RequestError{
		ErrorCode: status,
		Message:   message,
	}

	u.ServeJSON()
}
