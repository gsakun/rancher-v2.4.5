package client

const (
	CResourceType         = "cResource"
	CResourceFieldCpu     = "cpu"
	CResourceFieldGpu     = "gpu"
	CResourceFieldMemory  = "memory"
	CResourceFieldVolumes = "volumes"
)

type CResource struct {
	Cpu     string    `json:"cpu,omitempty" yaml:"cpu,omitempty"`
	Gpu     int64     `json:"gpu,omitempty" yaml:"gpu,omitempty"`
	Memory  string    `json:"memory,omitempty" yaml:"memory,omitempty"`
	Volumes []CVolume `json:"volumes,omitempty" yaml:"volumes,omitempty"`
}
