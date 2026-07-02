package model

type VerificationCodeRequest struct {
	Code string `json:"code" validate:"required,len=6,numeric"`
}

type VerificationCodeResponse struct {
	ID              string `json:"id"`
	EmailVerifiedAt string `json:"email_verified_at"`
}
