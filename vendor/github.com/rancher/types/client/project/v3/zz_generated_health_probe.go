package client

const (
	HealthProbeType                     = "healthProbe"
	HealthProbeFieldExec                = "exec"
	HealthProbeFieldFailureThreshold    = "failureThreshold"
	HealthProbeFieldHTTPGet             = "httpGet"
	HealthProbeFieldInitialDelaySeconds = "initialDelaySeconds"
	HealthProbeFieldPeriodSeconds       = "periodSeconds"
	HealthProbeFieldSuccessThreshold    = "successThreshold"
	HealthProbeFieldTCPSocket           = "tcpSocket"
	HealthProbeFieldTimeoutSeconds      = "timeoutSeconds"
)

type HealthProbe struct {
	Exec                *ExecAction      `json:"exec,omitempty" yaml:"exec,omitempty"`
	FailureThreshold    int64            `json:"failureThreshold,omitempty" yaml:"failureThreshold,omitempty"`
	HTTPGet             *HTTPGetAction   `json:"httpGet,omitempty" yaml:"httpGet,omitempty"`
	InitialDelaySeconds int64            `json:"initialDelaySeconds,omitempty" yaml:"initialDelaySeconds,omitempty"`
	PeriodSeconds       int64            `json:"periodSeconds,omitempty" yaml:"periodSeconds,omitempty"`
	SuccessThreshold    int64            `json:"successThreshold,omitempty" yaml:"successThreshold,omitempty"`
	TCPSocket           *TCPSocketAction `json:"tcpSocket,omitempty" yaml:"tcpSocket,omitempty"`
	TimeoutSeconds      int64            `json:"timeoutSeconds,omitempty" yaml:"timeoutSeconds,omitempty"`
}
