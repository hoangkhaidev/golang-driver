package req

import (
	"my-driver/model/validate_handle"

	"github.com/go-playground/validator/v10"
)
type ReqSignIn struct {
	Email 	 string `json:"email,omitempty" validate:"required,email" msg_required:"err:required_email" msg_email:"err:invalid_email"`
	Password string `json:"password,omitempty" validate:"required,gte=7" msg_required:"err:required_password" msg_gte:"err:password_gte_than_7_characters"`
}

func (m *ReqSignIn) Validate(validate *validator.Validate) error {
    return validate_handle.ValidateFunc[ReqSignIn](*m, validate)
}