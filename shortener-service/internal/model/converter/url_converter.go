package converter

import (
	"shortener-service/internal/entity"
	"shortener-service/internal/model"
	"shortener-service/internal/utils"
)

// =====================================================
// CREATE
// =====================================================

func ToCreateURLResponse(
	url *entity.URL,
) *model.CreateURLResponse {

	return &model.CreateURLResponse{
		ID:          url.ID,
		ShortCode:   url.ShortCode,
		OriginalURL: url.OriginalURL,
		IsActive:    url.IsActive,
		ClickCount:  url.ClickCount,
		ExpiredAt:   url.ExpiredAt,
		Passworded:  url.PasswordHash != nil,
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
// UPDATE
// =====================================================

func ToUpdateURLResponse(
	url *entity.URL,
) *model.UpdateURLResponse {

	return &model.UpdateURLResponse{
		ID:          url.ID,
		ShortCode:   url.ShortCode,
		OriginalURL: url.OriginalURL,
		IsActive:    url.IsActive,
		ClickCount:  url.ClickCount,
		ExpiredAt:   url.ExpiredAt,
		Passworded:  url.PasswordHash != nil,
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
// DETAIL
// =====================================================

func ToURLResponse(
	url *entity.URL,
) *model.URLResponse {

	return &model.URLResponse{
		ID:          url.ID,
		ShortCode:   url.ShortCode,
		OriginalURL: url.OriginalURL,
		IsActive:    url.IsActive,
		ClickCount:  url.ClickCount,
		ExpiredAt:   url.ExpiredAt,
		Passworded:  url.PasswordHash != nil,
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
// LIST
// =====================================================

func ToURLListResponse(
	url *entity.URL,
) *model.URLListResponse {

	return &model.URLListResponse{
		ID:          url.ID,
		ShortCode:   url.ShortCode,
		OriginalURL: url.OriginalURL,
		IsActive:    url.IsActive,
		ClickCount:  url.ClickCount,
		Passworded:  url.PasswordHash != nil,
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
// LIST MANY
// =====================================================

func ToURLListResponses(
	urls []*entity.URL,
) []*model.URLListResponse {

	responses := make(
		[]*model.URLListResponse,
		0,
		len(urls),
	)

	for _, url := range urls {

		responses = append(
			responses,
			ToURLListResponse(
				url,
			),
		)
	}

	return responses
}

// =====================================================
// DELETE
// =====================================================

func ToDeleteURLResponse(
	id string,
) *model.DeleteURLResponse {

	return &model.DeleteURLResponse{
		ID: id,
	}
}

// =====================================================
// RESTORE
// =====================================================

func ToRestoreURLResponse(
	id string,
) *model.RestoreURLResponse {

	return &model.RestoreURLResponse{
		ID: id,
	}
}

// =====================================================
// REDIRECT
// =====================================================

func ToRedirectResponse(
	originalURL string,
) *model.RedirectResponse {

	return &model.RedirectResponse{
		OriginalURL: originalURL,
	}
}
