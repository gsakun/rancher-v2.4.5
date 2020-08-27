package client

const (
	ApplicationSpecType            = "applicationSpec"
	ApplicationSpecFieldComponents = "components"
	ApplicationSpecFieldOptTraits  = "optTraits"
)

type ApplicationSpec struct {
	Components []Component            `json:"components,omitempty" yaml:"components,omitempty"`
	OptTraits  *ComponentTraitsForOpt `json:"optTraits,omitempty" yaml:"optTraits,omitempty"`
}
