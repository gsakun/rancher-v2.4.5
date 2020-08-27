package client

const (
	CEnvVarType           = "cEnvVar"
	CEnvVarFieldFromParam = "fromParam"
	CEnvVarFieldName      = "name"
	CEnvVarFieldValue     = "value"
)

type CEnvVar struct {
	FromParam string `json:"fromParam,omitempty" yaml:"fromParam,omitempty"`
	Name      string `json:"name,omitempty" yaml:"name,omitempty"`
	Value     string `json:"value,omitempty" yaml:"value,omitempty"`
}
