// link service
package service

import (
	"backend/api/dto"
	"backend/internal/dao"
)

type LinkService interface {
	CreateLink(req dto.CreateLinkRequest) (string, error)
	GetLink(shortCode string) (dto.GetLinkResponse, error)
	UpdateLink(shortCode string, req dto.UpdateLinkRequest) error
	DeleteLink(shortCode string) error
	ListLinks() ([]dto.GetLinkResponse, error)
}

type linkService struct {
	dao dao.LinkDAO
}

func NewLinkService(linkDAO dao.LinkDAO) LinkService {
	return &linkService{
		dao: linkDAO,
	}
}

// 创建短链接
func (s *linkService) CreateLink(req dto.CreateLinkRequest) (string, error) {
	return s.dao.Create(req)
}

// 获取短链接信息
func (s *linkService) GetLink(shortCode string) (dto.GetLinkResponse, error) {
	return s.dao.Get(shortCode)
}

// 更新短链接信息
func (s *linkService) UpdateLink(shortCode string, req dto.UpdateLinkRequest) error {
	return s.dao.Update(shortCode, req)
}

// 删除短链接
func (s *linkService) DeleteLink(shortCode string) error {
	err := s.dao.Delete(shortCode)
	if err != nil {
		return err
	}
	return nil
}

// 获取短链接列表
func (s *linkService) ListLinks() ([]dto.GetLinkResponse, error) {
	return s.dao.List()
}
