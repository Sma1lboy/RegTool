package source

type RegistryManager interface {
	GetCurrRegistry() (string, error)
	SetRegistry(string) (string, error)
}
