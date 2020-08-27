package client

const (
	ApplicationStatusType                   = "applicationStatus"
	ApplicationStatusFieldComponentResource = "componentResource"
)

type ApplicationStatus struct {
	ComponentResource map[string]ComponentResources `json:"componentResource,omitempty" yaml:"componentResource,omitempty"`
}
