package config

type Config interface {
	Load(content []byte) error

	Get(key string) interface{}
}

type ConfigAware interface {
	SetConfig(conf Config)
}
