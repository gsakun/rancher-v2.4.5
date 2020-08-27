package client

const (
	CustomMetricType        = "customMetric"
	CustomMetricFieldEnable = "enable"
	CustomMetricFieldUri    = "uri"
)

type CustomMetric struct {
	Enable bool   `json:"enable,omitempty" yaml:"enable,omitempty"`
	Uri    string `json:"uri,omitempty" yaml:"uri,omitempty"`
}
