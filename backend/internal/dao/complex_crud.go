package dao

import (
	"backend/api/dto"
	"backend/internal/dao/model"
	"errors"
	"gorm.io/gorm"
	"log"
	"time"
)

type LinkDAO interface {
	Create(req dto.CreateLinkRequest) (string, error)
	Get(shortCode string) (dto.GetLinkResponse, error)
	Update(shortCode string, req dto.UpdateLinkRequest) error
	Delete(shortCode string) error
	List() ([]dto.GetLinkResponse, error)
}

type linkDAO struct {
	db *gorm.DB
}

func NewLinkDAO(db *gorm.DB) LinkDAO {
	return &linkDAO{db: db}
}

// 添加短链接记录
func (dao *linkDAO) Create(req dto.CreateLinkRequest) (string, error) {
	startTime, err := time.Parse(time.RFC3339, req.StartTime)
	if err != nil {
		return "", err
	}
	endTime, err := time.Parse(time.RFC3339, req.EndTime)
	if err != nil {
		return "", err
	}

	// 检查 short_code 是否已经存在
	var existingLink model.Link
	if err := dao.db.Where("short_code = ?", req.ShortCode).First(&existingLink).Error; err == nil {
		return "", errors.New("short_code already exists")
	}

	link := model.Link{
		OriginalURL: req.OriginalURL,
		ShortCode:   req.ShortCode,
		StartTime:   startTime,
		EndTime:     endTime,
		IsActive:    true,
		Comment:     req.Comment,
	}

	if err := dao.db.Create(&link).Error; err != nil {
		return "", err
	}

	return link.ShortCode, nil
}

// 检索短链接记录
func (dao *linkDAO) Get(shortCode string) (dto.GetLinkResponse, error) {
	var link model.Link
	if err := dao.db.Where("short_code = ?", shortCode).First(&link).Error; err != nil {
		return dto.GetLinkResponse{}, err
	}
	return dto.GetLinkResponse{
		OriginalURL: link.OriginalURL,
		StartTime:   link.StartTime.String(),
		EndTime:     link.EndTime.String(),
		IsActive:    link.IsActive,
	}, nil
}

// 更新短链接记录
func (dao *linkDAO) Update(shortCode string, req dto.UpdateLinkRequest) error {
	var link model.Link
	if err := dao.db.Where("short_code = ?", shortCode).First(&link).Error; err != nil {
		return err
	}

	if req.OriginalURL != nil {
		link.OriginalURL = *req.OriginalURL
	}
	if req.StartTime != nil {
		startTime, err := time.Parse(time.RFC3339, *req.StartTime)
		if err != nil {
			return err
		}
		link.StartTime = startTime
	}
	if req.EndTime != nil {
		endTime, err := time.Parse(time.RFC3339, *req.EndTime)
		if err != nil {
			return err
		}
		link.EndTime = endTime
	}
	if req.IsActive != nil {
		link.IsActive = *req.IsActive
	}
	return dao.db.Save(&link).Error
}

// 删除短链接记录
func (dao *linkDAO) Delete(shortCode string) error {
	// 打印 short_code
	log.Printf("Attempting to delete link with short code: %s", shortCode)

	// 删除记录
	result := dao.db.Exec("DELETE FROM links WHERE short_code = ?", shortCode)

	// 打印消息
	if result.Error != nil {
		log.Printf("Error deleting link: %v", result.Error) // 删除失败
		return result.Error
	}

	if result.RowsAffected == 0 {
		log.Printf("No link found with short code: %s", shortCode) // 没有记录
		return errors.New("no link found with the specified short code")
	}

	log.Printf("Successfully deleted link with short code: %s", shortCode)
	return nil
}

// 获取所有短链接
func (dao *linkDAO) List() ([]dto.GetLinkResponse, error) {
	var links []model.Link
	if err := dao.db.Find(&links).Error; err != nil {
		return nil, err
	}

	var response []dto.GetLinkResponse
	for _, link := range links {
		response = append(response, dto.GetLinkResponse{
			OriginalURL: link.OriginalURL,
			StartTime:   link.StartTime.String(),
			EndTime:     link.EndTime.String(),
			IsActive:    link.IsActive,
		})
	}
	return response, nil
}
