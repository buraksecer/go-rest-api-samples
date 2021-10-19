package models

type CreateTokenRequest struct {
	// Username type is string and we validated
	Username string `form:"username" validate:"required_all=Password" qs:"username"`
	// Password type is string and we validated
	Password string `form:"password" validate:"required_all=Username" qs:"password"`
}
