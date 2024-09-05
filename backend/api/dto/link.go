package dto

type CreateLinkRequest struct {
	OriginalURL string `json:"origin" binding:"required"`
	ShortCode   string `json:"short"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Comment     string `json:"comment"`
}

type DeleteLinkRequest struct {
	ShortCode string `json:"short"`
}

type CreateLinkResponse struct {
	ShortCode string `json:"short_code"`
}

type GetLinkResponse struct {
	OriginalURL string `json:"original_url"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	IsActive    bool   `json:"is_active"`
}

type UpdateLinkRequest struct {
	OriginalURL *string `json:"original_url,omitempty"`
	StartTime   *string `json:"start_time,omitempty"`
	EndTime     *string `json:"end_time,omitempty"`
	IsActive    *bool   `json:"is_active,omitempty"`
}

type ListLinksResponse struct {
	Links []GetLinkResponse `json:"links"`
}
