package client

const (
	CircuitBreakingType                   = "circuitBreaking"
	CircuitBreakingFieldConnectionPool    = "connectionPool"
	CircuitBreakingFieldOutlierDetection  = "outlierDetection"
	CircuitBreakingFieldPortLevelSettings = "portLevelSettings"
)

type CircuitBreaking struct {
	ConnectionPool    *ConnectionPoolSettings `json:"connectionPool,omitempty" yaml:"connectionPool,omitempty"`
	OutlierDetection  *OutlierDetection       `json:"outlierDetection,omitempty" yaml:"outlierDetection,omitempty"`
	PortLevelSettings []PortTrafficPolicy     `json:"portLevelSettings,omitempty" yaml:"portLevelSettings,omitempty"`
}
