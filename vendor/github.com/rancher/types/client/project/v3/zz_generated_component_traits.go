package client

const (
	ComponentTraitsType                               = "componentTraits"
	ComponentTraitsFieldAutoscaling                   = "autoscaling"
	ComponentTraitsFieldCustomMetric                  = "custommetric"
	ComponentTraitsFieldLogcollect                    = "logcollect"
	ComponentTraitsFieldReplicas                      = "replicas"
	ComponentTraitsFieldSchedulePolicy                = "schedulePolicy"
	ComponentTraitsFieldTerminationGracePeriodSeconds = "terminationGracePeriodSeconds"
)

type ComponentTraits struct {
	Autoscaling                   *Autoscaling    `json:"autoscaling,omitempty" yaml:"autoscaling,omitempty"`
	CustomMetric                  *CustomMetric   `json:"custommetric,omitempty" yaml:"custommetric,omitempty"`
	Logcollect                    bool            `json:"logcollect,omitempty" yaml:"logcollect,omitempty"`
	Replicas                      int64           `json:"replicas,omitempty" yaml:"replicas,omitempty"`
	SchedulePolicy                *SchedulePolicy `json:"schedulePolicy,omitempty" yaml:"schedulePolicy,omitempty"`
	TerminationGracePeriodSeconds int64           `json:"terminationGracePeriodSeconds,omitempty" yaml:"terminationGracePeriodSeconds,omitempty"`
}
