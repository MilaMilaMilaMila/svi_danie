package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"svi_danie/internal/decorators"
	"svi_danie/internal/repositories/models"
	"svi_danie/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PageController struct {
	AuthDecorator *decorators.AuthDecorator

	PageService *services.PageService
	ProjService *services.ProjService
}

func (c *PageController) InitRouter(router gin.IRouter) {
	router.POST("/add_page", c.AuthDecorator.AddAuth(c.createPage))
	router.PUT("/edit_page", c.AuthDecorator.AddAuth(c.updatePage))
	router.DELETE("/delete_page", c.AuthDecorator.AddAuth(c.deletePage))
	router.GET("/get_page", c.AuthDecorator.AddAuth(c.getPage))
	router.GET("/get_all_pages", c.AuthDecorator.AddAuth(c.getAllProjectPages))
}

func (c *PageController) createPage(ctx *gin.Context, user *models.User) {
	type Response struct {
		PageId uuid.UUID `json:"page_id"`
	}

	createdPage, err := validateCreatePageRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = c.ProjService.CheckOwnership(user.Id, createdPage.ProjectId)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	if err = c.PageService.CreatePage(createdPage); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Response{PageId: createdPage.Id})
}

func (c *PageController) updatePage(ctx *gin.Context, user *models.User) {
	type Response struct {
		PageId uuid.UUID `json:"page_id"`
	}

	page, err := validateUpdatePageRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = c.PageService.GetPage(page.Id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = c.ProjService.CheckOwnership(user.Id, page.ProjectId)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	if err = c.PageService.UpdatePage(page); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, Response{PageId: page.Id})
}

func (c *PageController) deletePage(ctx *gin.Context, user *models.User) {
	type Response struct {
		PageId uuid.UUID `json:"page_id"`
	}

	pageId, err := uuid.Parse(ctx.Query("page_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingPage, err := c.PageService.GetPage(pageId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = c.ProjService.CheckOwnership(user.Id, existingPage.ProjectId)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	if err = c.PageService.DeletePage(pageId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Response{PageId: pageId})
}

func (c *PageController) getPage(ctx *gin.Context, user *models.User) {
	type Response struct {
		Page *models.Page `json:"page"`
	}

	pageId, err := uuid.Parse(ctx.Query("page_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	page, err := c.PageService.GetPage(pageId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = c.ProjService.CheckOwnership(user.Id, page.ProjectId); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, Response{Page: page})
}

func (c *PageController) getAllProjectPages(ctx *gin.Context, user *models.User) {
	type Response struct {
		Pages []*models.Page `json:"pages"`
	}

	projectId, err := uuid.Parse(ctx.Query("project_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.ProjService.CheckOwnership(user.Id, projectId)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	pages, err := c.PageService.GetAllProjectPages(projectId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, Response{Pages: pages})
}

func validateCreatePageRequest(ctx *gin.Context) (*models.Page, error) {
	projectId, err := uuid.Parse(ctx.Query("project_id"))
	if err != nil {
		return nil, err
	}

	title := ctx.Query("title")

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return nil, err
	}

	var data json.RawMessage
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &models.Page{
		Id:        uuid.New(),
		ProjectId: projectId,
		Title:     title,
		Data:      data,
	}, nil
}

func validateUpdatePageRequest(ctx *gin.Context) (*models.Page, error) {
	pageId, err := uuid.Parse(ctx.Query("id"))
	if err != nil {
		return nil, err
	}

	projectId, err := uuid.Parse(ctx.Query("project_id"))
	if err != nil {
		return nil, err
	}

	title := ctx.Query("title")

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return nil, err
	}

	var data json.RawMessage
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &models.Page{
		Id:        pageId,
		ProjectId: projectId,
		Title:     title,
		Data:      data,
	}, nil
}
