package source

import "registryhub/source/structs"

type RegistryManager interface {
	GetCurrRegistry() (string, error)
	SetRegistry(structs.Region, *structs.RegistrySources) (string, error)
}
