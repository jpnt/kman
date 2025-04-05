package core

type IStep interface {
	Name() string
	// NeedParams() []string // TODO: simpler way?
	// This serves for the use case of when I only want
	// to run a set of steps. And for those steps if I have
	// not defined the required kernel context parameters
	// then it should warn the users. Speficy in the command
	// line arguments or prompt the user? What is more suckless
	Execute() error
}
