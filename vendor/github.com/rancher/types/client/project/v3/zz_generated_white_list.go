package client

const (
	WhiteListType       = "whiteList"
	WhiteListFieldUsers = "users"
)

type WhiteList struct {
	Users []string `json:"users,omitempty" yaml:"users,omitempty"`
}
