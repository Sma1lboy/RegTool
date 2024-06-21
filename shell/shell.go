package shell

type Shell interface {
	SetEnv(key, value string) error
}
