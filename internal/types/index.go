package types

type LoginPayload struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type RegisterPayload struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type AuthData struct {
	User  any
	Token any
}
