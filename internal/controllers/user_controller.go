package controllers

import (
	"errors"
	"net/http"
	"svi_danie/internal/repositories/models"
	"svi_danie/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	UserService *services.UserService
}

func (c *UserController) InitRouter(router gin.IRouter) {
	router.POST("/add_user", c.createUser)
	router.GET("/get_user", c.getUser)
}

func (c *UserController) createUser(ctx *gin.Context) {
	type Response struct {
		UserId uuid.UUID `json:"user_id"`
	}

	user := &models.User{
		Id:       uuid.New(),
		Login:    ctx.Query("login"),
		Password: ctx.Query("password"),
	}

	err := c.UserService.AddUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, Response{UserId: user.Id})
}

func (c *UserController) getUser(ctx *gin.Context) {
	type Response struct {
		Id    uuid.UUID `json:"id"`
		Login string    `json:"login"`
	}

	user, err := c.authUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, Response{Login: user.Login, Id: user.Id})
}

func (d *UserController) authUser(ctx *gin.Context) (*models.User, error) {
	login, password, ok := ctx.Request.BasicAuth()
	if !ok {
		return nil, errors.New("no authentication provided")
	}
	user, err := d.UserService.AuthUser(login, password)
	if err != nil {
		return nil, errors.New("bad authentication provided")
	}
	return user, nil
}
