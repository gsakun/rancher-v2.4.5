package client

const (
	ImagePullConfigType          = "imagePullConfig"
	ImagePullConfigFieldPassword = "password"
	ImagePullConfigFieldRegistry = "registry"
	ImagePullConfigFieldUsername = "username"
)

type ImagePullConfig struct {
	Password string `json:"password,omitempty" yaml:"password,omitempty"`
	Registry string `json:"registry,omitempty" yaml:"registry,omitempty"`
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}
