package model

// =====================================================
// QR CODE
// =====================================================

type GetQRCodeRequest struct {
	ID string `params:"id" validate:"required,uuid"`
}
