package client

const (
	ConfigFileType           = "configFile"
	ConfigFileFieldFileName  = "fileName"
	ConfigFileFieldFromParam = "fromParam"
	ConfigFileFieldPath      = "path"
	ConfigFileFieldValue     = "value"
)

type ConfigFile struct {
	FileName  string `json:"fileName,omitempty" yaml:"fileName,omitempty"`
	FromParam string `json:"fromParam,omitempty" yaml:"fromParam,omitempty"`
	Path      string `json:"path,omitempty" yaml:"path,omitempty"`
	Value     string `json:"value,omitempty" yaml:"value,omitempty"`
}
