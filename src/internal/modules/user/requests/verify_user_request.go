package user_requests

type VerifyUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
