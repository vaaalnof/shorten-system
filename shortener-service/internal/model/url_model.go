package model

// =====================================================
// CREATE URL
// =====================================================

type AddURLRequest struct {
	OriginalURL string `json:"original_url" validate:"required,max=2048,shorturl"`
	ShortCode   string `json:"short_code" validate:"omitempty,min=3,max=32,shortcode"`
	Password    string `json:"password" validate:"omitempty,min=7,max=100"`
	ExpiredAt   int64  `json:"expired_at"`
}

type AddURLResponse struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
}

// =====================================================
// GET URL
// =====================================================

type GetURLRequest struct {
	ID string `params:"id" validate:"required,uuid"`
}

type URLResponse struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	ShortCode   string `json:"short_code"`
	OriginalURL string `json:"original_url"`
	IsActive    bool   `json:"is_active"`
	HasPassword bool   `json:"has_password"`
	ExpiredAt   *int64 `json:"expired_at,omitempty"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// =====================================================
// CACHE
// =====================================================

type CachedURL struct {
	ID           string  `json:"id"`
	ShortCode    string  `json:"short_code"`
	OriginalURL  string  `json:"original_url"`
	IsActive     bool    `json:"is_active"`
	PasswordHash *string `json:"password_hash,omitempty"`
	ExpiredAt    *int64  `json:"expired_at,omitempty"`
}

// =====================================================
// UPDATE PASSWORD
// =====================================================

type UpdateURLPasswordRequest struct {
	ID          string `params:"id" validate:"required,uuid"`
	OldPassword string `json:"old_password" validate:"omitempty,min=7,max=100"`
	NewPassword string `json:"new_password" validate:"required,min=7,max=100"`
}

type UpdateURLPasswordResponse struct {
	ID        string `json:"id"`
	UpdatedAt string `json:"updated_at"`
}

// =====================================================
// REMOVE PASSWORD
// =====================================================

type RemoveURLPasswordRequest struct {
	ID          string `params:"id" validate:"required,uuid"`
	OldPassword string `json:"old_password" validate:"required,min=7,max=100"`
}

// =====================================================
// VERIFY PASSWORD
// =====================================================

type VerifyURLPasswordRequest struct {
	ShortCode string `params:"short_code" validate:"required,min=3,max=32,shortcode"`
	Password  string `json:"password" validate:"required,min=7,max=100"`
}

// =====================================================
// REDIRECT URL
// =====================================================

type RedirectURLRequest struct {
	ShortCode string `params:"short_code" validate:"required,min=3,max=32,shortcode"`
}

// =====================================================
// UPDATE STATUS
// =====================================================

type UpdateURLStatusRequest struct {
	ID       string `params:"id" validate:"required,uuid"`
	IsActive bool   `json:"is_active"`
}

type UpdateURLStatusResponse struct {
	ID        string `json:"id"`
	IsActive  bool   `json:"is_active"`
	UpdatedAt string `json:"updated_at"`
}

// =====================================================
// LIST URLS
// =====================================================

type ListURLsRequest struct {
	Page int `query:"page"`
	Size int `query:"size"`
}

type ListURLResponse struct {
	ID          string `json:"id"`
	ShortCode   string `json:"short_code"`
	OriginalURL string `json:"original_url"`
	IsActive    bool   `json:"is_active"`
	HasPassword bool   `json:"has_password"`
	ExpiredAt   *int64 `json:"expired_at,omitempty"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// =====================================================
// DELETE URL
// =====================================================

type DeleteURLRequest struct {
	ID string `params:"id" validate:"required,uuid"`
}

type DeleteURLResponse struct {
	ID        string `json:"id"`
	UpdatedAt string `json:"updated_at"`
}

// =====================================================
// UPDATE URL
// =====================================================

type UpdateURLRequest struct {
	ID          string  `params:"id" validate:"required,uuid"`
	OriginalURL *string `json:"original_url,omitempty" validate:"omitempty,max=2048,shorturl"`
	ShortCode   *string `json:"short_code,omitempty" validate:"omitempty,min=3,max=32,shortcode"`
	IsActive    *bool   `json:"is_active,omitempty"`
	ExpiredAt   *int64  `json:"expired_at,omitempty"`
}

type UpdateURLResponse struct {
	ID        string `json:"id"`
	UpdatedAt string `json:"updated_at"`
}
