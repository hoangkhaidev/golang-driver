package banana

import "errors"

var (
	WrongPassword = errors.New("Wrong password!")
	UserNotFound = errors.New("User not found!")
	UserConflict = errors.New("User already exists!")
	SignUpFail = errors.New("Register fail!")
	EmailRequired = errors.New("Email required!")
)