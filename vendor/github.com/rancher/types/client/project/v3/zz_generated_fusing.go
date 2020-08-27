package client

const (
	FusingType         = "fusing"
	FusingFieldAction  = "action"
	FusingFieldPodList = "podlist"
)

type Fusing struct {
	Action  string   `json:"action,omitempty" yaml:"action,omitempty"`
	PodList []string `json:"podlist,omitempty" yaml:"podlist,omitempty"`
}
