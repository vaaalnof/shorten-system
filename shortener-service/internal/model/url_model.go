package model

// =====================================================
// CREATE URL
// =====================================================

type CreateURLRequest struct {
	OriginalURL string  `json:"original_url" validate:"required,url"`
	ShortCode   *string `json:"short_code,omitempty" validate:"omitempty,min=3,max=32"`
	Password    *string `json:"password,omitempty" validate:"omitempty,min=6,max=100"`
	ExpiredAt   *int64  `json:"expired_at,omitempty"`
}

type CreateURLResponse struct {
	ID          string `json:"id"`
	ShortCode   string `json:"short_code"`
	OriginalURL string `json:"original_url"`
	IsActive    bool   `json:"is_active"`
	ClickCount  int64  `json:"click_count"`
	ExpiredAt   *int64 `json:"expired_at,omitempty"`
	Passworded  bool   `json:"passworded"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// =====================================================
// UPDATE URL
// =====================================================

type UpdateURLRequest struct {
	OriginalURL *string `json:"original_url,omitempty" validate:"omitempty,url"`
	IsActive    *bool   `json:"is_active,omitempty"`
	Password    *string `json:"password,omitempty" validate:"omitempty,min=6,max=100"`
	ExpiredAt   *int64  `json:"expired_at,omitempty"`
}

type UpdateURLResponse struct {
	ID          string `json:"id"`
	ShortCode   string `json:"short_code"`
	OriginalURL string `json:"original_url"`
	IsActive    bool   `json:"is_active"`
	ClickCount  int64  `json:"click_count"`
	ExpiredAt   *int64 `json:"expired_at,omitempty"`
	Passworded  bool   `json:"passworded"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// =====================================================
// GET URL
// =====================================================

type URLResponse struct {
	ID          string `json:"id"`
	ShortCode   string `json:"short_code"`
	OriginalURL string `json:"original_url"`
	IsActive    bool   `json:"is_active"`
	ClickCount  int64  `json:"click_count"`
	ExpiredAt   *int64 `json:"expired_at,omitempty"`
	Passworded  bool   `json:"passworded"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// =====================================================
// LIST URL
// =====================================================

type URLListResponse struct {
	ID          string `json:"id"`
	ShortCode   string `json:"short_code"`
	OriginalURL string `json:"original_url"`
	IsActive    bool   `json:"is_active"`
	ClickCount  int64  `json:"click_count"`
	Passworded  bool   `json:"passworded"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// =====================================================
// DELETE URL
// =====================================================

type DeleteURLResponse struct {
	ID string `json:"id"`
}

// =====================================================
// RESTORE URL
// =====================================================

type RestoreURLResponse struct {
	ID string `json:"id"`
}

// =====================================================
// REDIRECT
// =====================================================

type RedirectResponse struct {
	OriginalURL string `json:"original_url"`
}
