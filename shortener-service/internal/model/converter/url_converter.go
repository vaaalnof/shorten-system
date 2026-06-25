package converter

import (
	"shortener-service/internal/entity"
	"shortener-service/internal/model"
	"shortener-service/internal/utils"
)

// =====================================================
// CREATE URL
// =====================================================

func ToAddURLResponse(
	url *entity.URL,
) *model.AddURLResponse {

	return &model.AddURLResponse{
		ID: url.ID,
		CreatedAt: utils.FormatUnixTime(
			url.CreatedAt,
			"2006-01-02 15:04:05",
		),
	}
}

// =====================================================
// GET URL
// =====================================================

func ToURLResponse(
	url *entity.URL,
) *model.URLResponse {

	return &model.URLResponse{
		ID:          url.ID,
		UserID:      url.UserID,
		ShortCode:   url.ShortCode,
		OriginalURL: url.OriginalURL,
		IsActive:    url.IsActive,
		HasPassword: url.PasswordHash != nil,
		ExpiredAt:   url.ExpiredAt,
		CreatedAt: utils.FormatUnixTime(
			url.CreatedAt,
			"2006-01-02 15:04:05",
		),
		UpdatedAt: utils.FormatUnixTime(
			url.UpdatedAt,
			"2006-01-02 15:04:05",
		),
	}
}

// =====================================================
// UPDATE PASSWORD
// =====================================================

func ToUpdateURLPasswordResponse(
	url *entity.URL,
) *model.UpdateURLPasswordResponse {

	return &model.UpdateURLPasswordResponse{
		ID: url.ID,
		UpdatedAt: utils.FormatUnixTime(
			url.UpdatedAt,
			"2006-01-02 15:04:05",
		),
	}

}

// =====================================================
// UPDATE STATUS
// =====================================================

func ToUpdateURLStatusResponse(
	url *entity.URL,
) *model.UpdateURLStatusResponse {

	return &model.UpdateURLStatusResponse{
		ID:       url.ID,
		IsActive: url.IsActive,
		UpdatedAt: utils.FormatUnixTime(
			url.UpdatedAt,
			"2006-01-02 15:04:05",
		),
	}
}

// =====================================================
// LIST URLS
// =====================================================

func ToListURLResponse(
	url *entity.URL,
) *model.ListURLResponse {

	return &model.ListURLResponse{
		ID:          url.ID,
		ShortCode:   url.ShortCode,
		OriginalURL: url.OriginalURL,
		IsActive:    url.IsActive,
		HasPassword: url.PasswordHash != nil,
		ExpiredAt:   url.ExpiredAt,
		CreatedAt: utils.FormatUnixTime(
			url.CreatedAt,
			"2006-01-02 15:04:05",
		),
		UpdatedAt: utils.FormatUnixTime(
			url.UpdatedAt,
			"2006-01-02 15:04:05",
		),
	}
}

func ToListURLResponses(
	urls []*entity.URL,
) []*model.ListURLResponse {

	responses := make(
		[]*model.ListURLResponse,
		0,
		len(urls),
	)

	for _, url := range urls {

		responses = append(
			responses,
			ToListURLResponse(url),
		)
	}

	return responses
}

// =====================================================
// DELETE URL
// =====================================================

func ToDeleteURLResponse(
	url *entity.URL,
) *model.DeleteURLResponse {

	return &model.DeleteURLResponse{
		ID: url.ID,
		UpdatedAt: utils.FormatUnixTime(
			url.UpdatedAt,
			"2006-01-02 15:04:05",
		),
	}
}

// =====================================================
// UPDATE URL
// =====================================================

func ToUpdateURLResponse(
	url *entity.URL,
) *model.UpdateURLResponse {

	return &model.UpdateURLResponse{
		ID: url.ID,
		UpdatedAt: utils.FormatUnixTime(
			url.UpdatedAt,
			"2006-01-02 15:04:05",
		),
	}
}
