package core

type IStep interface {
	Name() string
	// CtxParamsNeeded() []string // TODO: simpler way?
	Execute() error
	// Completed() bool
}
