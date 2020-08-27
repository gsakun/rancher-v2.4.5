package client

const (
	ParameterType             = "parameter"
	ParameterFieldDefault     = "default"
	ParameterFieldDescription = "description"
	ParameterFieldName        = "name"
	ParameterFieldRequired    = "required"
	ParameterFieldType        = "type"
)

type Parameter struct {
	Default     string `json:"default,omitempty" yaml:"default,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Name        string `json:"name,omitempty" yaml:"name,omitempty"`
	Required    bool   `json:"required,omitempty" yaml:"required,omitempty"`
	Type        string `json:"type,omitempty" yaml:"type,omitempty"`
}
