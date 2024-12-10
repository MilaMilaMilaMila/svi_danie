package controllers

import (
	"net/http"
	"svi_danie/internal/decorators"
	"svi_danie/internal/repositories/models"
	"svi_danie/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProjectController struct {
	AuthDecorator *decorators.AuthDecorator

	ProjectService *services.ProjService
}

func (c *ProjectController) InitRouter(router gin.IRouter) {
	router.POST("/add_proj", c.AuthDecorator.AddAuth(c.createProject))
	router.DELETE("/delete_proj", c.AuthDecorator.AddAuth(c.deleteProject))
	router.GET("/get_proj", c.AuthDecorator.AddAuth(c.getProject))
	router.GET("/get_all_proj", c.AuthDecorator.AddAuth(c.getAllProjects))
}

func (c *ProjectController) createProject(ctx *gin.Context, user *models.User) {
	type Response struct {
		ProjId uuid.UUID `json:"proj_id"`
	}

	title := ctx.Query("title")

	project := &models.Project{
		Id:      uuid.New(),
		OwnerId: user.Id,
		Title:   title,
	}

	if err := c.ProjectService.CreateProj(project); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, Response{ProjId: project.Id})
}

func (c *ProjectController) deleteProject(ctx *gin.Context, user *models.User) {
	type Response struct {
		ProjId uuid.UUID `json:"proj_id"`
	}

	projectId, err := uuid.Parse(ctx.Query("project_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.ProjectService.CheckOwnership(user.Id, projectId)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	err = c.ProjectService.DeleteProj(projectId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Response{ProjId: projectId})
}

func (c *ProjectController) getProject(ctx *gin.Context, user *models.User) {
	type Response struct {
		Proj *models.Project `json:"proj"`
	}

	projectId, err := uuid.Parse(ctx.Query("project_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project, err := c.ProjectService.GetProj(projectId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if project.OwnerId != user.Id {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "project ownership does not match"})
		return
	}

	ctx.JSON(http.StatusOK, Response{Proj: project})
}

func (c *ProjectController) getAllProjects(ctx *gin.Context, user *models.User) {
	type Response struct {
		Projects []*models.Project `json:"projects"`
	}

	projects, err := c.ProjectService.GetAllUserProj(user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Response{Projects: projects})
}
