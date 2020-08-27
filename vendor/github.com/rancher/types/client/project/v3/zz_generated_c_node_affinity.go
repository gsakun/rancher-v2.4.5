package client

const (
	CNodeAffinityType                           = "cNodeAffinity"
	CNodeAffinityFieldCLabelSelectorRequirement = "labelSelectorRequirement"
	CNodeAffinityFieldHardAffinity              = "hardAffinity"
)

type CNodeAffinity struct {
	CLabelSelectorRequirement *CLabelSelectorRequirement `json:"labelSelectorRequirement,omitempty" yaml:"labelSelectorRequirement,omitempty"`
	HardAffinity              bool                       `json:"hardAffinity,omitempty" yaml:"hardAffinity,omitempty"`
}
