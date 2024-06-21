package source

import "registryhub/source/structs"

// RegistryManager is an interface for managing registries.
// This interface defines methods to get the current registry URL and set the registry URL based on a specified region and sources.
type RegistryManager interface {

	// GetCurrRegistry retrieves the current registry URL.
	// Returns:
	// - string: The current registry URL.
	// - error: An error if there is an issue retrieving the URL.
	GetCurrRegistry() (string, error)

	// SetRegistry sets the registry URL based on the specified region and sources.
	// This method needs to be implemented differently depending on the operating system.
	// Parameters:
	// - region: The specified region for which the registry URL should be set.
	// - sources: A pointer to a RegistrySources struct containing mappings of regions to their respective URLs.
	// Returns:
	// - string: A message indicating the result of the operation.
	// - error: An error if there is an issue setting the registry URL.
	SetRegistry(region structs.Region, sources *structs.RegistrySources) (string, error)
}
