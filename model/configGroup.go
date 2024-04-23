package model

type ConfigGroup struct {
	Name    string            `json:"name"`
	Version int               `json:"version"`
	Configs map[string]Config `json:"configs"`
}

type ConfigGroupRepository interface {
	Get(name string, version int) (ConfigGroup, error)
	Add(c ConfigGroup)
	Delete(name string, version int) error
	AddConfigToGroup(group ConfigGroup, config Config) error
	RemoveConfigFromGroup(group ConfigGroup, key string) error
}
