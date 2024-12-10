package controllers

import (
	"io"
	"net/http"
	"svi_danie/internal/decorators"
	"svi_danie/internal/repositories/models"
	"svi_danie/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ImageController struct {
	AuthDecorator *decorators.AuthDecorator

	ImgService *services.ImgService
}

func (c *ImageController) InitRouter(router gin.IRouter) {
	router.POST("/create_img", c.AuthDecorator.AddAuth(c.createImage))
	router.GET("/get_img", c.AuthDecorator.AddAuth(c.GetImage))
}

func (c *ImageController) createImage(ctx *gin.Context, _ *models.User) {
	type Response struct {
		ImageUrl string `json:"img_url"`
	}

	img, err := validateCreateImageRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = c.ImgService.CreateImage(img); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, formUrl(img.Id, ctx.Request))
}

func (c *ImageController) GetImage(ctx *gin.Context, _ *models.User) {
	imageId, err := uuid.Parse(ctx.Param("img_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	img, err := c.ImgService.GetImageById(imageId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.Writer.Header().Set("Content-Type", http.DetectContentType(img.Data))
	ctx.Writer.WriteHeader(http.StatusOK)
	_, err = ctx.Writer.Write(img.Data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func formUrl(imgId uuid.UUID, r *http.Request) string {
	protocol := r.URL.Scheme
	if protocol == "" {
		if r.TLS != nil {
			protocol = "https"
		} else {
			protocol = "http"
		}
	}
	return protocol + "://" + r.Host + "/get_img?img_id=" + imgId.String()
}

func validateCreateImageRequest(ctx *gin.Context) (*models.Img, error) {
	err := ctx.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		return nil, err
	}

	imgFile, _, err := ctx.Request.FormFile("img")
	if err != nil {
		return nil, err
	}

	imgBytes, err := io.ReadAll(imgFile)
	if err != nil {
		return nil, err
	}

	return &models.Img{
		Id:   uuid.New(),
		Data: imgBytes,
	}, nil
}
