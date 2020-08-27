package client

const (
	AppPortType               = "appPort"
	AppPortFieldContainerPort = "containerPort"
	AppPortFieldName          = "name"
	AppPortFieldProtocol      = "protocol"
)

type AppPort struct {
	ContainerPort int64  `json:"containerPort,omitempty" yaml:"containerPort,omitempty"`
	Name          string `json:"name,omitempty" yaml:"name,omitempty"`
	Protocol      string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
}
