package request

type SignInReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type SignUpReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
	Username string `json:"username" validate:"required,min=2"`
}

type ReIssueReq struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type UpdatePasswordReq struct {
	AlreadyPassword string `json:"alreadyPassword" validate:"required,password"`
	NewPassword     string `json:"newPassword" validate:"required,password"`
}
