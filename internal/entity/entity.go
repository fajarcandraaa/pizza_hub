package entity

type Error string

// Declare error messege
const (
	ErrPermissionNotAllowed = Error("permission.not_allowed")
	ErrExpiredTime          = Error("Expired Link URL")
	ErrStillInProgressTime  = Error("Previous process has not been completed")

	//User Error
	ErrUserNotExist             = Error("domain.user.error.not_exist")
	ErrUserAlreadyExist         = Error("domain.user.error.email_or_username_alredy_exist")
	ErrUsersCredentialNotExist  = Error("domain.user.error.credential_not_exist")
	ErrUsersUnprocessableEntity = Error("domain.user.error.unprocessable_entity")

	//Menu Error
	ErrMenuNotExist             = Error("domain.Menu.error.not_exist")
	ErrMenuAlreadyExist         = Error("domain.Menu.error.menu_alredy_exist")
	ErrMenusUnprocessableEntity = Error("domain.user.error.unprocessable_entity")
)

func (e Error) Error() string {
	return string(e)
}
