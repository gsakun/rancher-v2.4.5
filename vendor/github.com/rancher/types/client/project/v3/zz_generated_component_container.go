package client

const (
	ComponentContainerType                 = "componentContainer"
	ComponentContainerFieldArgs            = "args"
	ComponentContainerFieldCommand         = "command"
	ComponentContainerFieldConfig          = "config"
	ComponentContainerFieldEnv             = "env"
	ComponentContainerFieldImage           = "image"
	ComponentContainerFieldImagePullPolicy = "imagePullPolicy"
	ComponentContainerFieldImagePullSecret = "imagePullSecret"
	ComponentContainerFieldLifecycle       = "lifecycle"
	ComponentContainerFieldLivenessProbe   = "livenessProbe"
	ComponentContainerFieldName            = "name"
	ComponentContainerFieldPorts           = "ports"
	ComponentContainerFieldReadinessProbe  = "readinessProbe"
	ComponentContainerFieldResources       = "resources"
	ComponentContainerFieldSecurityContext = "securityContext"
)

type ComponentContainer struct {
	Args            []string         `json:"args,omitempty" yaml:"args,omitempty"`
	Command         []string         `json:"command,omitempty" yaml:"command,omitempty"`
	Config          []ConfigFile     `json:"config,omitempty" yaml:"config,omitempty"`
	Env             []CEnvVar        `json:"env,omitempty" yaml:"env,omitempty"`
	Image           string           `json:"image,omitempty" yaml:"image,omitempty"`
	ImagePullPolicy string           `json:"imagePullPolicy,omitempty" yaml:"imagePullPolicy,omitempty"`
	ImagePullSecret string           `json:"imagePullSecret,omitempty" yaml:"imagePullSecret,omitempty"`
	Lifecycle       *CLifecycle      `json:"lifecycle,omitempty" yaml:"lifecycle,omitempty"`
	LivenessProbe   *HealthProbe     `json:"livenessProbe,omitempty" yaml:"livenessProbe,omitempty"`
	Name            string           `json:"name,omitempty" yaml:"name,omitempty"`
	Ports           []AppPort        `json:"ports,omitempty" yaml:"ports,omitempty"`
	ReadinessProbe  *HealthProbe     `json:"readinessProbe,omitempty" yaml:"readinessProbe,omitempty"`
	Resources       *CResource       `json:"resources,omitempty" yaml:"resources,omitempty"`
	SecurityContext *SecurityContext `json:"securityContext,omitempty" yaml:"securityContext,omitempty"`
}
