package alias

// AliasManager manages primary names and their aliases
type AliasManager struct {
	primaryToAliases map[string][]string
	aliasToPrimary   map[string]string
}

// newAliasManager creates a new AliasManager
func newAliasManager() *AliasManager {
	return &AliasManager{
		primaryToAliases: make(map[string][]string),
		aliasToPrimary:   make(map[string]string),
	}
}

// manager is the global instance of AliasManager
var manager = newAliasManager()

// RegisterAlias registers a primary name with its aliases
func RegisterAlias(primary string, aliases []string) {
	manager.primaryToAliases[primary] = aliases
	manager.aliasToPrimary[primary] = primary
	for _, alias := range aliases {
		manager.aliasToPrimary[alias] = primary
	}
}

// GetPrimary returns the primary name for a given alias
func GetPrimary(alias string) string {
	if primary, ok := manager.aliasToPrimary[alias]; ok {
		return primary
	}
	return alias
}

// GetAllAliases returns all aliases for a given primary name
func GetAllAliases(primary string) []string {
	if aliases, ok := manager.primaryToAliases[primary]; ok {
		return aliases
	}
	return []string{}
}
func GetAllPrimary() []string {
	res := make([]string, 0)
	for k := range manager.primaryToAliases {
		res = append(res, k)
	}
	return res
}
