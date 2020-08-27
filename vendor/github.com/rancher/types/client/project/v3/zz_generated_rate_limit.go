package client

const (
	RateLimitType               = "rateLimit"
	RateLimitFieldOverrides     = "overrides"
	RateLimitFieldRequestAmount = "requestAmount"
	RateLimitFieldTimeDuration  = "timeDuration"
)

type RateLimit struct {
	Overrides     []Override `json:"overrides,omitempty" yaml:"overrides,omitempty"`
	RequestAmount int64      `json:"requestAmount,omitempty" yaml:"requestAmount,omitempty"`
	TimeDuration  string     `json:"timeDuration,omitempty" yaml:"timeDuration,omitempty"`
}
