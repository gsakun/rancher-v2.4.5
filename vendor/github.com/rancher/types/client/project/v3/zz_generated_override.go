package client

const (
	OverrideType               = "override"
	OverrideFieldRequestAmount = "requestAmount"
	OverrideFieldUser          = "user"
)

type Override struct {
	RequestAmount int64  `json:"requestAmount,omitempty" yaml:"requestAmount,omitempty"`
	User          string `json:"user,omitempty" yaml:"user,omitempty"`
}
