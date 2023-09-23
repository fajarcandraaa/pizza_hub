package entity

type Error string

// Declare error messege
const (
	ErrPermissionNotAllowed = Error("permission.not_allowed")
	ErrExpiredTime          = Error("Expired Link URL")

	//User Error
	ErrUserNotExist             = Error("domain.user.error.not_exist")
	ErrUserAlreadyExist         = Error("domain.user.error.email_or_username_alredy_exist")
	ErrUsersCredentialNotExist  = Error("domain.user.error.credential_not_exist")
	ErrUsersUnprocessableEntity = Error("domain.user.error.unprocessable_entity")
)

func (e Error) Error() string {
	return string(e)
}
