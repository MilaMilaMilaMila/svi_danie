package decorators

import (
	"errors"
	"net/http"
	"svi_danie/internal/repositories/models"
	"svi_danie/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthDecorator struct {
	UserService *services.UserService
}

func (d *AuthDecorator) AddAuth(handler func(*gin.Context, *models.User)) func(*gin.Context) {
	return func(ctx *gin.Context) {
		user, err := d.authUser(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		handler(ctx, user)
	}
}

func (d *AuthDecorator) authUser(ctx *gin.Context) (*models.User, error) {
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
