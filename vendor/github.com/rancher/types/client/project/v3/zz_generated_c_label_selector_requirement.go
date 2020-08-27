package client

const (
	CLabelSelectorRequirementType          = "cLabelSelectorRequirement"
	CLabelSelectorRequirementFieldKey      = "key"
	CLabelSelectorRequirementFieldOperator = "operator"
	CLabelSelectorRequirementFieldValues   = "values"
)

type CLabelSelectorRequirement struct {
	Key      string   `json:"key,omitempty" yaml:"key,omitempty"`
	Operator string   `json:"operator,omitempty" yaml:"operator,omitempty"`
	Values   []string `json:"values,omitempty" yaml:"values,omitempty"`
}
