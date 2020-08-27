package client

const (
	ComponentResourcesType                    = "componentResources"
	ComponentResourcesFieldClusterRbacConfig  = "clusterRbacConfig"
	ComponentResourcesFieldComponentId        = "componentId"
	ComponentResourcesFieldConfigMaps         = "configMaps"
	ComponentResourcesFieldDestinationRule    = "DestinationRule"
	ComponentResourcesFieldGateway            = "gateway"
	ComponentResourcesFieldImagePullSecret    = "imagePullSecret"
	ComponentResourcesFieldPolicy             = "policy"
	ComponentResourcesFieldService            = "service"
	ComponentResourcesFieldServiceRole        = "serviceRole"
	ComponentResourcesFieldServiceRoleBinding = "serviceRoleBinding"
	ComponentResourcesFieldVirtualService     = "virtualService"
	ComponentResourcesFieldWorkload           = "workload"
)

type ComponentResources struct {
	ClusterRbacConfig  string   `json:"clusterRbacConfig,omitempty" yaml:"clusterRbacConfig,omitempty"`
	ComponentId        string   `json:"componentId,omitempty" yaml:"componentId,omitempty"`
	ConfigMaps         []string `json:"configMaps,omitempty" yaml:"configMaps,omitempty"`
	DestinationRule    string   `json:"DestinationRule,omitempty" yaml:"DestinationRule,omitempty"`
	Gateway            string   `json:"gateway,omitempty" yaml:"gateway,omitempty"`
	ImagePullSecret    string   `json:"imagePullSecret,omitempty" yaml:"imagePullSecret,omitempty"`
	Policy             string   `json:"policy,omitempty" yaml:"policy,omitempty"`
	Service            string   `json:"service,omitempty" yaml:"service,omitempty"`
	ServiceRole        string   `json:"serviceRole,omitempty" yaml:"serviceRole,omitempty"`
	ServiceRoleBinding string   `json:"serviceRoleBinding,omitempty" yaml:"serviceRoleBinding,omitempty"`
	VirtualService     string   `json:"virtualService,omitempty" yaml:"virtualService,omitempty"`
	Workload           string   `json:"workload,omitempty" yaml:"workload,omitempty"`
}
