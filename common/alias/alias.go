package alias

// AliasManager manages primary names and their aliases
type AliasManager struct {
	aliases map[string]string
}

// NewAliasManager creates a new AliasManager
func NewAliasManager() *AliasManager {
	return &AliasManager{
		aliases: make(map[string]string),
	}
}

// RegisterAlias registers a primary name with its aliases
func (am *AliasManager) RegisterAlias(primary string, aliases []string) {
	am.aliases[primary] = primary
	for _, alias := range aliases {
		am.aliases[alias] = primary
	}
}

// GetPrimary returns the primary name for a given alias
func (am *AliasManager) GetPrimary(alias string) string {
	if primary, ok := am.aliases[alias]; ok {
		return primary
	}
	return alias
}

// GetAllAliases returns all aliases for a given primary name
func (am *AliasManager) GetAllAliases(primary string) []string {
	var result []string
	for alias, p := range am.aliases {
		if p == primary {
			result = append(result, alias)
		}
	}
	return result
}
