package client

const (
	ComponentType                  = "component"
	ComponentFieldArch             = "arch"
	ComponentFieldComponentTraits  = "componentTraits"
	ComponentFieldContainers       = "containers"
	ComponentFieldName             = "name"
	ComponentFieldOsType           = "osType"
	ComponentFieldParameters       = "parameters"
	ComponentFieldVersion          = "version"
	ComponentFieldWorkloadSettings = "workloadSetings"
	ComponentFieldWorkloadType     = "workloadType"
)

type Component struct {
	Arch             string               `json:"arch,omitempty" yaml:"arch,omitempty"`
	ComponentTraits  *ComponentTraits     `json:"componentTraits,omitempty" yaml:"componentTraits,omitempty"`
	Containers       []ComponentContainer `json:"containers,omitempty" yaml:"containers,omitempty"`
	Name             string               `json:"name,omitempty" yaml:"name,omitempty"`
	OsType           string               `json:"osType,omitempty" yaml:"osType,omitempty"`
	Parameters       []Parameter          `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Version          string               `json:"version,omitempty" yaml:"version,omitempty"`
	WorkloadSettings []WorkloadSetting    `json:"workloadSetings,omitempty" yaml:"workloadSetings,omitempty"`
	WorkloadType     string               `json:"workloadType,omitempty" yaml:"workloadType,omitempty"`
}
