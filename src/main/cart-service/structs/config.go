package structs

type Config struct {
	Port      int64  `yaml:"port"`
	ApiPrefix string `yaml:"apiPrefix"`
	DbAddress string `yaml:"dbAddress"`
}
