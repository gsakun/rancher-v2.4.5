package client

const (
	VolumeMounterType              = "volumeMounter"
	VolumeMounterFieldStorageClass = "storageClass"
	VolumeMounterFieldVolumeName   = "volumeName"
)

type VolumeMounter struct {
	StorageClass string `json:"storageClass,omitempty" yaml:"storageClass,omitempty"`
	VolumeName   string `json:"volumeName,omitempty" yaml:"volumeName,omitempty"`
}
