package userhandler

type UserDto struct {
	Id       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
} // @name UserDto

type CreateUserRequest struct {
	Username string `json:"username"`
	IdToken  string `json:"idToken"`
} // @name CreateUserRequest

type CheckUserByIdTokenRequest struct {
	IdToken string `json:"idToken"`
} // @name CheckUserByIdTokenRequest

type CheckUserByIdTokenResponse struct {
	Exist bool `json:"exist"`
} // @name CheckUserByIdTokenResponse
