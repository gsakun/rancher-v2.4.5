package client

const (
	WorkloadSettingType           = "workloadSetting"
	WorkloadSettingFieldFromParam = "fromParam"
	WorkloadSettingFieldName      = "name"
	WorkloadSettingFieldType      = "type"
	WorkloadSettingFieldValue     = "value"
)

type WorkloadSetting struct {
	FromParam string `json:"fromParam,omitempty" yaml:"fromParam,omitempty"`
	Name      string `json:"name,omitempty" yaml:"name,omitempty"`
	Type      string `json:"type,omitempty" yaml:"type,omitempty"`
	Value     string `json:"value,omitempty" yaml:"value,omitempty"`
}
