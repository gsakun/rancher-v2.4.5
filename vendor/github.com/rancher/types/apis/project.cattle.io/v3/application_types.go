package v3

import (
	"github.com/rancher/norman/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// PullAlways means that kubelet always attempts to pull the latest image. Container will fail If the pull fails.
	PullAlways PullPolicy = "Always"
	// PullNever means that kubelet never pulls an image, but only uses a local image. Container will fail if the image isn't present
	PullNever PullPolicy = "Never"
	// PullIfNotPresent means that kubelet pulls if the image isn't present on disk. Container will fail if the image isn't present and the pull fails.
	PullIfNotPresent PullPolicy = "IfNotPresent"
)

const (
	Server          WorkloadType = "Server"
	SingletonServer WorkloadType = "SingletonServer"
	Worker          WorkloadType = "Worker"
	SingletonWorker WorkloadType = "SingletonWorker"
	Task            WorkloadType = "Task"
	SingletonTask   WorkloadType = "SingletonTaskTask"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Application is a specification for a Application resource

type Application struct {
	types.Namespaced
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApplicationSpec   `json:"spec,omitempty"`
	Status ApplicationStatus `json:"status,omitempty"`
}

type ApplicationSpec struct {
	Components []Component `json:"components"`
	//DevTraits  ComponentTraitsForDev `json:"devTraits,omitempty"`
	OptTraits ComponentTraitsForOpt `json:"optTraits,omitempty"`
}

type WhiteList struct {
	Users []string `json:"users,omitempty"`
}

type AppIngress struct {
	Host       string `json:"host"`
	Path       string `json:"path,omitempty"`
	ServerPort int32  `json:"serverPort"`
}

type VolumeMounter struct {
	VolumeName   string `json:"volumeName"`
	StorageClass string `json:"storageClass"`
}

type ComponentTraitsForOpt struct {
	LoadBalancer    *LoadBalancerSettings `json:"loadBalancer,omitempty"`
	GrayRelease     map[string]int        `json:"grayRelease,omitempty"`
	ImagePullConfig *ImagePullConfig      `json:"imagePullConfig,omitempty"`
	StaticIP        bool                  `json:"staticIP,omitempty"`
	VolumeMounter   *VolumeMounter        `json:"volumeMounter,omitempty"`
	Ingress         AppIngress            `json:"ingress"`
	WhiteList       *WhiteList            `json:"whiteList,omitempty"`
	Eject           []string              `json:"eject,omitempty"`
	Fusing          *Fusing               `json:"fusing,omitempty"` //zk
	RateLimit       *RateLimit            `json:"rateLimit,omitempty"`
	CircuitBreaking *CircuitBreaking      `json:"circuitbreaking,omitempty"` //zk
	HTTPRetry       *HTTPRetry            `json:"httpretry,omitempty"`
}

//zk
type CustomMetric struct {
	Enable bool   `json:"enable"`
	Uri    string `json:"uri,omitempty"`
}

type Autoscaling struct {
	Metric      string `json:"metric"`
	Threshold   int32  `json:"threshold"`
	MaxReplicas int32  `json:"maxreplicas"`
	MinReplicas int32  `json:"minreplicas"`
}

type HTTPRetry struct {
	Attempts      int    `json:"attempts"`
	PerTryTimeout string `json:"perTryTimeout"`
}

//zk
type CircuitBreaking struct {
	ConnectionPool    *ConnectionPoolSettings `json:"connectionPool,omitempty"`
	OutlierDetection  *OutlierDetection       `json:"outlierDetection,omitempty"`
	PortLevelSettings []PortTrafficPolicy     `json:"portLevelSettings,omitempty"`
}

type ConnectionPoolSettings struct {

	// Settings common to both HTTP and TCP upstream connections.
	TCP *TCPSettings `json:"tcp,omitempty"`

	// HTTP connection pool settings.
	HTTP *HTTPSettings `json:"http,omitempty"`
}

type PortSelector struct {
	// Choose one of the fields below.

	// Valid port number
	Number uint32 `json:"number,omitempty"`

	// Valid port name
	Name string `json:"name,omitempty"`
}

type OutlierDetection struct {
	ConsecutiveErrors  int32  `json:"consecutiveErrors,omitempty"`
	Interval           string `json:"interval,omitempty"`
	BaseEjectionTime   string `json:"baseEjectionTime,omitempty"`
	MaxEjectionPercent int32  `json:"maxEjectionPercent,omitempty"`
}

// Settings common to both HTTP and TCP upstream connections.
type TCPSettings struct {
	MaxConnections int32 `json:"maxConnections,omitempty"`

	ConnectTimeout string `json:"connectTimeout,omitempty"`
}

// Settings applicable to HTTP1.1/HTTP2/GRPC connections.
type HTTPSettings struct {
	HTTP1MaxPendingRequests  int32 `json:"http1MaxPendingRequests,omitempty"`
	HTTP2MaxRequests         int32 `json:"http2MaxRequests,omitempty"`
	MaxRequestsPerConnection int32 `json:"maxRequestsPerConnection,omitempty"`
	MaxRetries               int32 `json:"maxRetries,omitempty"`
}

type SimpleLB string

const (
	SimpleLBRoundRobin  SimpleLB = "ROUND_ROBIN"
	SimpleLBLeastConn   SimpleLB = "LEAST_CONN"
	SimpleLBRandom      SimpleLB = "RANDOM"
	SimpleLBPassthrough SimpleLB = "PASSTHROUGH"
)

//zk
type Fusing struct {
	PodList []string `json:"podlist,omitempty"`
	Action  string   `json:"action,omitempty"`
}

type RateLimit struct {
	TimeDuration  string     `json:"timeDuration"`
	RequestAmount int32      `json:"requestAmount"`
	Overrides     []Override `json:"overrides,omitempty"`
}

type Override struct {
	RequestAmount int32  `json:"requestAmount"`
	User          string `json:"user"`
}

// Traffic policies that apply to specific ports of the service
type PortTrafficPolicy struct {
	Port             PortSelector           `json:"port"`
	LoadBalancer     LoadBalancerSettings   `json:"loadBalancer,omitempty"`
	ConnectionPool   ConnectionPoolSettings `json:"connectionPool,omitempty"`
	OutlierDetection OutlierDetection       `json:"outlierDetection,omitempty"`
}
type LoadBalancerSettings struct {
	Simple         SimpleLB          `json:"simple,omitempty"`
	ConsistentHash *ConsistentHashLB `json:"consistentHash,omitempty"`
}

type ConsistentHashLB struct {
	HTTPHeaderName  string `json:"httpHeaderName,omitempty"`
	UseSourceIP     bool   `json:"useSourceIp,omitempty"`
	MinimumRingSize uint64 `json:"minimumRingSize,omitempty"`
}

//负载均衡类型 rr;leastConn;random
//consistentType sourceIP
type IngressLB struct {
	LBType         string `json:"lbType,omitempty"`
	ConsistentType string `json:"consistentType,omitempty"`
}

type ImagePullConfig struct {
	Registry string `json:"registry,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type ComponentTraits struct {
	Replicas                      int32           `json:"replicas"`
	CustomMetric                  *CustomMetric   `json:"custommetric,omitempty"` //zk
	Logcollect                    bool            `json:"logcollect,,omitempty"`
	TerminationGracePeriodSeconds int64           `json:"terminationGracePeriodSeconds,omitempty"` //zk
	SchedulePolicy                *SchedulePolicy `json:"schedulePolicy,omitempty"`
	Autoscaling                   *Autoscaling    `json:"autoscaling,omitempty"` //zk
}

type Disk struct {
	Required  string `json:"required,omitempty"`
	Ephemeral bool   `json:"ephemeral"`
}

type CVolume struct {
	Name          string `json:"name"`
	MountPath     string `json:"mountPath"`
	AccessMode    string `json:"accessMode,omitempty"`
	SharingPolicy string `json:"sharingPolicy,omitempty"`
	Disk          Disk   `json:"disk"`
}

type CResource struct {
	Cpu     string    `json:"cpu,omitempty"`
	Memory  string    `json:"memory,omitempty"`
	Gpu     int       `json:"gpu,omitempty"`
	Volumes []CVolume `json:"volumes,omitempty"`
}

type CEnvVar struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	FromParam string `json:"fromParam,omitempty"`
}

type AppPort struct {
	Name          string `json:"name,omitempty"`
	ContainerPort int32  `json:"containerPort"`
	Protocol      string `json:"protocol,omitempty"`
}

type ComponentContainer struct {
	Name string `json:"name"`

	Image string `json:"image"`

	Command []string `json:"command,omitempty"`

	Args []string `json:"args,omitempty"`

	Ports []AppPort `json:"ports,omitempty"`

	Env []CEnvVar `json:"env,omitempty"`

	Resources CResource `json:"resources,omitempty"`

	LivenessProbe *HealthProbe `json:"livenessProbe,omitempty"`

	ReadinessProbe *HealthProbe `json:"readinessProbe,omitempty"`

	ImagePullPolicy PullPolicy       `json:"imagePullPolicy,omitempty"`
	Lifecycle       *CLifecycle      `json:"lifecycle,omitempty"`
	Config          []ConfigFile     `json:"config,omitempty"`
	ImagePullSecret string           `json:"imagePullSecret,omitempty"`
	SecurityContext *SecurityContext `json:"securityContext,omitempty"`
}

type SchedulePolicy struct {
	NodeSelector    map[string]string `json:"nodeSelector,omitempty"`
	NodeAffinity    *CNodeAffinity    `json:"nodeAffinity,omitempty"`
	PodAffinity     *CPodAffinity     `json:"podAffinity,omitempty"`
	PodAntiAffinity *CPodAntiAffinity `json:"podAntiAffinity,omitempty"`
}

type CPodAntiAffinity struct {
	HardAffinity               bool `json:"hardAffinity,omitempty"`
	*CLabelSelectorRequirement `json:"labelSelectorRequirement,omitempty"`
}

type CPodAffinity struct {
	HardAffinity               bool `json:"hardAffinity,omitempty"`
	*CLabelSelectorRequirement `json:"labelSelectorRequirement,omitempty"`
}

type CNodeAffinity struct {
	HardAffinity               bool `json:"hardAffinity,omitempty"`
	*CLabelSelectorRequirement `json:"labelSelectorRequirement,omitempty"`
}

type CLabelSelectorRequirement struct {
	Key      string                 `json:"key" patchStrategy:"merge" patchMergeKey:"key" protobuf:"bytes,1,opt,name=key"`
	Operator CLabelSelectorOperator `json:"operator" protobuf:"bytes,2,opt,name=operator,casttype=LabelSelectorOperator"`
	Values   []string               `json:"values,omitempty" protobuf:"bytes,3,rep,name=values"`
}

// A label selector operator is the set of operators that can be used in a selector requirement.
type CLabelSelectorOperator string

const (
	LabelSelectorOpIn           CLabelSelectorOperator = "In"
	LabelSelectorOpNotIn        CLabelSelectorOperator = "NotIn"
	LabelSelectorOpExists       CLabelSelectorOperator = "Exists"
	LabelSelectorOpDoesNotExist CLabelSelectorOperator = "DoesNotExist"
)

type CLifecycle struct {
	PostStart *Handler `json:"postStart,omitempty" protobuf:"bytes,1,opt,name=postStart"`
	PreStop   *Handler `json:"preStop,omitempty" protobuf:"bytes,2,opt,name=preStop"`
}

type WorkloadType string

type Component struct {
	Name         string       `json:"name"`
	Version      string       `json:"version"`
	Parameters   []Parameter  `json:"parameters,omitempty"`
	WorkloadType WorkloadType `json:"workloadType"`

	OsType string `json:"osType,omitempty"`

	Arch string `json:"arch,omitempty"`

	Containers       []ComponentContainer `json:"containers,omitempty"`
	ComponentTraits  ComponentTraits      `json:"componentTraits,omitempty"`
	WorkloadSettings []WorkloadSetting    `json:"workloadSetings,omitempty"`
}

//int,float,string,bool,json
type Parameter struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type"`
	Required    bool   `json:"required,omitempty"`
	Default     string `json:"default,omitempty"`
}

type SecurityContext struct{}

type ConfigFile struct {
	Path      string `json:"path"`
	FileName  string `json:"fileName"`
	Value     string `json:"value"`
	FromParam string `json:"fromParam,omitempty"`
}

type WorkloadSetting struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Value     string `json:"value"`
	FromParam string `json:"fromParam"`
}

type Handler struct {
	Exec      *ExecAction      `json:"exec,omitempty" protobuf:"bytes,1,opt,name=exec"`
	HTTPGet   *HTTPGetAction   `json:"httpGet,omitempty" protobuf:"bytes,2,opt,name=httpGet"`
	TCPSocket *TCPSocketAction `json:"tcpSocket,omitempty" protobuf:"bytes,3,opt,name=tcpSocket"`
}

type HealthProbe struct {
	Handler             `json:haneler",inline" protobuf:"bytes,1,opt,name=handler"`
	InitialDelaySeconds int32 `json:"initialDelaySeconds,omitempty" protobuf:"varint,2,opt,name=initialDelaySeconds"`

	TimeoutSeconds int32 `json:"timeoutSeconds,omitempty" protobuf:"varint,3,opt,name=timeoutSeconds"`

	PeriodSeconds int32 `json:"periodSeconds,omitempty" protobuf:"varint,4,opt,name=periodSeconds"`

	SuccessThreshold int32 `json:"successThreshold,omitempty" protobuf:"varint,5,opt,name=successThreshold"`

	FailureThreshold int32 `json:"failureThreshold,omitempty" protobuf:"varint,6,opt,name=failureThreshold"`
}

type TCPSocketAction struct {
	// Number or name of the port to access on the container.
	// Number must be in the range 1 to 65535.
	// Name must be an IANA_SVC_NAME.
	Port int `json:"port" protobuf:"bytes,1,opt,name=port"`
}

type HTTPGetAction struct {
	// Path to access on the HTTP server.
	// +optional
	Path string `json:"path,omitempty" protobuf:"bytes,1,opt,name=path"`
	// Name or number of the port to access on the container.
	// Number must be in the range 1 to 65535.
	// Name must be an IANA_SVC_NAME.
	Port int `json:"port" protobuf:"bytes,2,opt,name=port"`
	// Host name to connect to, defaults to the pod IP. You probably want to set
	// "Host" in httpHeaders instead.
	// +optional

	HTTPHeaders []HTTPHeader `json:"httpHeaders,omitempty" protobuf:"bytes,5,rep,name=httpHeaders"`
}

type HTTPHeader struct {
	// The header field name
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
	// The header field value
	Value string `json:"value" protobuf:"bytes,2,opt,name=value"`
}

type ExecAction struct {
	Command []string `json:"command,omitempty" protobuf:"bytes,1,rep,name=command"`
}

type PullPolicy string

type ApplicationStatus struct {
	ComponentResource map[string]ComponentResources `json:"componentResource,omitempty"`
}

type ComponentResources struct {
	ComponentId        string   `json:"componentId,omitempty"`
	Workload           string   `json:"workload,omitempty"`
	Service            string   `json:"service,omitempty"`
	ConfigMaps         []string `json:"configMaps,omitempty"`
	ImagePullSecret    string   `json:"imagePullSecret,omitempty"`
	Gateway            string   `json:"gateway,omitempty"`
	Policy             string   `json:"policy,omitempty"`
	ClusterRbacConfig  string   `json:"clusterRbacConfig,omitempty"`
	VirtualService     string   `json:"virtualService,omitempty"`
	ServiceRole        string   `json:"serviceRole,omitempty"`
	ServiceRoleBinding string   `json:"serviceRoleBinding,omitempty"`
	DestinationRule    string   `json:"DestinationRule,omitempty"`
}
