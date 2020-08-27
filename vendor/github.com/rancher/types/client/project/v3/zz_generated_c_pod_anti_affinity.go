package client

const (
	CPodAntiAffinityType                           = "cPodAntiAffinity"
	CPodAntiAffinityFieldCLabelSelectorRequirement = "labelSelectorRequirement"
	CPodAntiAffinityFieldHardAffinity              = "hardAffinity"
)

type CPodAntiAffinity struct {
	CLabelSelectorRequirement *CLabelSelectorRequirement `json:"labelSelectorRequirement,omitempty" yaml:"labelSelectorRequirement,omitempty"`
	HardAffinity              bool                       `json:"hardAffinity,omitempty" yaml:"hardAffinity,omitempty"`
}
