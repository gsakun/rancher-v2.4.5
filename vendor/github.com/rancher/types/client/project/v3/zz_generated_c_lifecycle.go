package client

const (
	CLifecycleType           = "cLifecycle"
	CLifecycleFieldPostStart = "postStart"
	CLifecycleFieldPreStop   = "preStop"
)

type CLifecycle struct {
	PostStart *Handler `json:"postStart,omitempty" yaml:"postStart,omitempty"`
	PreStop   *Handler `json:"preStop,omitempty" yaml:"preStop,omitempty"`
}
