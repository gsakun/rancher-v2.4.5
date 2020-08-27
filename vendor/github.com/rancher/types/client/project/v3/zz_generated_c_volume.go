package client

const (
	CVolumeType               = "cVolume"
	CVolumeFieldAccessMode    = "accessMode"
	CVolumeFieldDisk          = "disk"
	CVolumeFieldMountPath     = "mountPath"
	CVolumeFieldName          = "name"
	CVolumeFieldSharingPolicy = "sharingPolicy"
)

type CVolume struct {
	AccessMode    string `json:"accessMode,omitempty" yaml:"accessMode,omitempty"`
	Disk          *Disk  `json:"disk,omitempty" yaml:"disk,omitempty"`
	MountPath     string `json:"mountPath,omitempty" yaml:"mountPath,omitempty"`
	Name          string `json:"name,omitempty" yaml:"name,omitempty"`
	SharingPolicy string `json:"sharingPolicy,omitempty" yaml:"sharingPolicy,omitempty"`
}
