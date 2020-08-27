package client

const (
	DiskType           = "disk"
	DiskFieldEphemeral = "ephemeral"
	DiskFieldRequired  = "required"
)

type Disk struct {
	Ephemeral bool   `json:"ephemeral,omitempty" yaml:"ephemeral,omitempty"`
	Required  string `json:"required,omitempty" yaml:"required,omitempty"`
}
