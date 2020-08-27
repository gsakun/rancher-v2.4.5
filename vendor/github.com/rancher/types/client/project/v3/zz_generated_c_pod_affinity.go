package client

const (
	CPodAffinityType                           = "cPodAffinity"
	CPodAffinityFieldCLabelSelectorRequirement = "labelSelectorRequirement"
	CPodAffinityFieldHardAffinity              = "hardAffinity"
)

type CPodAffinity struct {
	CLabelSelectorRequirement *CLabelSelectorRequirement `json:"labelSelectorRequirement,omitempty" yaml:"labelSelectorRequirement,omitempty"`
	HardAffinity              bool                       `json:"hardAffinity,omitempty" yaml:"hardAffinity,omitempty"`
}
