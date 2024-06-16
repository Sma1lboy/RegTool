package common

// interface that defines the methods to get and set the new registry
type RegistryManager interface {
	GetCurrRegistry() (string, error)
	SetRegistry(string) (string, error)
}
