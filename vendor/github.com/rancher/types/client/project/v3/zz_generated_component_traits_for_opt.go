package client

const (
	ComponentTraitsForOptType                 = "componentTraitsForOpt"
	ComponentTraitsForOptFieldCircuitBreaking = "circuitbreaking"
	ComponentTraitsForOptFieldEject           = "eject"
	ComponentTraitsForOptFieldFusing          = "fusing"
	ComponentTraitsForOptFieldGrayRelease     = "grayRelease"
	ComponentTraitsForOptFieldHTTPRetry       = "httpretry"
	ComponentTraitsForOptFieldImagePullConfig = "imagePullConfig"
	ComponentTraitsForOptFieldIngress         = "ingress"
	ComponentTraitsForOptFieldLoadBalancer    = "loadBalancer"
	ComponentTraitsForOptFieldRateLimit       = "rateLimit"
	ComponentTraitsForOptFieldStaticIP        = "staticIP"
	ComponentTraitsForOptFieldVolumeMounter   = "volumeMounter"
	ComponentTraitsForOptFieldWhiteList       = "whiteList"
)

type ComponentTraitsForOpt struct {
	CircuitBreaking *CircuitBreaking      `json:"circuitbreaking,omitempty" yaml:"circuitbreaking,omitempty"`
	Eject           []string              `json:"eject,omitempty" yaml:"eject,omitempty"`
	Fusing          *Fusing               `json:"fusing,omitempty" yaml:"fusing,omitempty"`
	GrayRelease     map[string]int64      `json:"grayRelease,omitempty" yaml:"grayRelease,omitempty"`
	HTTPRetry       *HTTPRetry            `json:"httpretry,omitempty" yaml:"httpretry,omitempty"`
	ImagePullConfig *ImagePullConfig      `json:"imagePullConfig,omitempty" yaml:"imagePullConfig,omitempty"`
	Ingress         *AppIngress           `json:"ingress,omitempty" yaml:"ingress,omitempty"`
	LoadBalancer    *LoadBalancerSettings `json:"loadBalancer,omitempty" yaml:"loadBalancer,omitempty"`
	RateLimit       *RateLimit            `json:"rateLimit,omitempty" yaml:"rateLimit,omitempty"`
	StaticIP        bool                  `json:"staticIP,omitempty" yaml:"staticIP,omitempty"`
	VolumeMounter   *VolumeMounter        `json:"volumeMounter,omitempty" yaml:"volumeMounter,omitempty"`
	WhiteList       *WhiteList            `json:"whiteList,omitempty" yaml:"whiteList,omitempty"`
}
