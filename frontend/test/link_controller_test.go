package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/api/dto"
	"backend/internal/controller"
	"backend/internal/dao"
	"backend/internal/service"
	"backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}
	return db
}

func setupRouter(linkCtrl *controller.LinkController) *gin.Engine {
	r := gin.Default()
	r.POST("/api/link/create", linkCtrl.CreateLink)
	r.POST("/api/link/delete", linkCtrl.DeleteLink)
	r.GET("/api/link/:short_code", linkCtrl.GetLink)
	return r
}

func TestCreateLink(t *testing.T) {
	db := setupTestDB(t)
	linkDAO := dao.NewLinkDAO(db)
	linkService := service.NewLinkService(linkDAO)
	linkCtrl := controller.NewLinkController(linkService)
	router := setupRouter(linkCtrl)

	linkRequest := dto.CreateLinkRequest{
		OriginalURL: "https://example.com",
		ShortCode:   "exmpl",
		StartTime:   "2024-09-05T00:00:00Z",
		EndTime:     "2024-12-31T00:00:00Z",
		Comment:     "Test link",
	}
	body, _ := json.Marshal(linkRequest)

	req, _ := http.NewRequest("POST", "/api/link/create", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "short_code")
}

func TestDeleteLink(t *testing.T) {
	db := setupTestDB(t)
	linkDAO := dao.NewLinkDAO(db)
	linkService := service.NewLinkService(linkDAO)
	linkCtrl := controller.NewLinkController(linkService)
	router := setupRouter(linkCtrl)

	link := dao.NewLinkDAO(db)
	link.Create(dto.CreateLinkRequest{
		OriginalURL: "https://example.com",
		ShortCode:   "testlink",
		StartTime:   "2024-09-05T00:00:00Z",
		EndTime:     "2024-12-31T00:00:00Z",
		Comment:     "Test link",
	})

	deleteRequest := dto.DeleteLinkRequest{
		ShortCode: "testlink",
	}
	body, _ := json.Marshal(deleteRequest)
	req, _ := http.NewRequest("POST", "/api/link/delete", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Link deleted successfully")
}
