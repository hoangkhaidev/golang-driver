package req

import (
	"my-driver/model/validate_handle"

	"github.com/go-playground/validator/v10"
)
type ReqUpdateProfile struct {
	FullName string `json:"full_name,omitempty" validate:"required" msg_required:"err:required_full_name"`
	Email 	 string `json:"email,omitempty" validate:"required,email" msg_required:"err:required_email" msg_email:"err:invalid_email"`
}

func (m *ReqUpdateProfile) Validate(validate *validator.Validate) error {
    return validate_handle.ValidateFunc[ReqUpdateProfile](*m, validate)
}