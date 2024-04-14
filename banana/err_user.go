package banana

import "errors"

var (
	WrongPasswordText = errors.New("Wrong password!")
	UserNotFoundText = errors.New("User not found!")
	UserConflictText = errors.New("User already exists!")
	SignUpFailText = errors.New("Register fail!")
	EmailRequiredText = errors.New("Email required!")
	UserNotUpdate = errors.New("User not update!")


	ErrUserNotFound = errors.New("err:user_not_found")
	ErrRequiredFullName = errors.New("err:required_full_name")
	ErrRequiredEmail = errors.New("err:required_email")
	ErrInvalidEmail = errors.New("err:invalid_email")
	ErrRequiredPassword = errors.New("err:required_password")
	ErrWrongPassword = errors.New("err:wrong_password")
	ErrPasswordGteThan7Characters = errors.New("err:password_gte_than_7_characters")
)