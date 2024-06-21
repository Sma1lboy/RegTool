package structs

type Region string

const (
	CN Region = "cn"
	US Region = "us"
	EU Region = "eu"
)

func StringToRegion(region string) Region {
	switch region {
	case "cn":
		return CN
	case "us":
		return US
	case "eu":
		return EU
	default:
		return ""
	}
}

// RegistrySources is a map of regions to registry regions
type RegistrySources map[Region]RegistryRegionSources

// RegistryRegionSources is a map of package managers to urls
type RegistryRegionSources map[string][]string
