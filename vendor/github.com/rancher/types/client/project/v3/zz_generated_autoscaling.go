package client

const (
	AutoscalingType             = "autoscaling"
	AutoscalingFieldMaxReplicas = "maxreplicas"
	AutoscalingFieldMetric      = "metric"
	AutoscalingFieldMinReplicas = "minreplicas"
	AutoscalingFieldThreshold   = "threshold"
)

type Autoscaling struct {
	MaxReplicas int64  `json:"maxreplicas,omitempty" yaml:"maxreplicas,omitempty"`
	Metric      string `json:"metric,omitempty" yaml:"metric,omitempty"`
	MinReplicas int64  `json:"minreplicas,omitempty" yaml:"minreplicas,omitempty"`
	Threshold   int64  `json:"threshold,omitempty" yaml:"threshold,omitempty"`
}
