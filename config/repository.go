package config

type Repository interface {
	Load(content []byte) error

	Get(key string) interface{}

	GetSub(key string) Repository
}
