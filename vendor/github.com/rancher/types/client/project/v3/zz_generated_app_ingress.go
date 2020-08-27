package client

const (
	AppIngressType            = "appIngress"
	AppIngressFieldHost       = "host"
	AppIngressFieldPath       = "path"
	AppIngressFieldServerPort = "serverPort"
)

type AppIngress struct {
	Host       string `json:"host,omitempty" yaml:"host,omitempty"`
	Path       string `json:"path,omitempty" yaml:"path,omitempty"`
	ServerPort int64  `json:"serverPort,omitempty" yaml:"serverPort,omitempty"`
}
