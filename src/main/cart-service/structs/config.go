package structs

type Config struct {
	Port     uint64 `yaml:"port,omitempty"`
	Instance uint64 `yaml:"instance,omitempty"`
}
