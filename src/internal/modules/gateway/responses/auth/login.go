package gateway_response

type Login struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
