package shellenv

type Shell interface {
	SetEnv(key, value string) error
}
