package client

const (
	SchedulePolicyType                 = "schedulePolicy"
	SchedulePolicyFieldNodeAffinity    = "nodeAffinity"
	SchedulePolicyFieldNodeSelector    = "nodeSelector"
	SchedulePolicyFieldPodAffinity     = "podAffinity"
	SchedulePolicyFieldPodAntiAffinity = "podAntiAffinity"
)

type SchedulePolicy struct {
	NodeAffinity    *CNodeAffinity    `json:"nodeAffinity,omitempty" yaml:"nodeAffinity,omitempty"`
	NodeSelector    map[string]string `json:"nodeSelector,omitempty" yaml:"nodeSelector,omitempty"`
	PodAffinity     *CPodAffinity     `json:"podAffinity,omitempty" yaml:"podAffinity,omitempty"`
	PodAntiAffinity *CPodAntiAffinity `json:"podAntiAffinity,omitempty" yaml:"podAntiAffinity,omitempty"`
}
