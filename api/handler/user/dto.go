package userhandler

type UserDto struct {
	Id       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
} // @name UserDto

type CreateUserRequest struct {
	Username string `json:"username"`
	IdToken  string `json:"idToken"`
} // @name CreateUserRequest
