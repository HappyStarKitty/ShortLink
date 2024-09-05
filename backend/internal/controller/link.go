// link controller
package controller

import (
	"backend/api/dto"
	"backend/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LinkController interface {
	CreateLink(c *gin.Context)
	GetLink(c *gin.Context)
	UpdateLink(c *gin.Context)
	DeleteLink(c *gin.Context)
	ListLinks(c *gin.Context)
}

type linkController struct {
	service service.LinkService
}

func NewLinkController(service service.LinkService) LinkController {
	return &linkController{
		service: service,
	}
}

// 创建短链接
func (ctrl *linkController) CreateLink(c *gin.Context) {
	var req dto.CreateLinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	shortCode, err := ctrl.service.CreateLink(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"short_code": shortCode})
}

// 获取短链接信息
func (ctrl *linkController) GetLink(c *gin.Context) {
	shortCode := c.Param("shortCode")
	link, err := ctrl.service.GetLink(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusMovedPermanently, link.OriginalURL)
}

// 更新短链接信息
func (ctrl *linkController) UpdateLink(c *gin.Context) {
	var req dto.UpdateLinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	shortCode := c.Param("shortCode")
	err := ctrl.service.UpdateLink(shortCode, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Link updated successfully"})
}

// 删除短链接
func (ctrl *linkController) DeleteLink(c *gin.Context) {
	var req dto.DeleteLinkRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	shortCode := req.ShortCode

	err := ctrl.service.DeleteLink(shortCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Link deleted successfully"})
}

// 获取短链接列表
func (ctrl *linkController) ListLinks(c *gin.Context) {
	links, err := ctrl.service.ListLinks() // No arguments passed
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"links": links})
}
