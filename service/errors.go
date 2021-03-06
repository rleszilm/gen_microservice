package service

import "errors"

var (
	// ErrDependencyCycle is returned when services canot be started because of a
	// cycle in dependencies.
	ErrDependencyCycle = errors.New("dependency cycle")

	// ErrMissingDependency is returned when a service requires a dependency that is
	// not registered.
	ErrMissingDependency = errors.New("missing dependency")

	// ErrNoProxy is returned if a server tries to attack a proxy that isn't configured.
	ErrNoProxy = errors.New("missing proxy config")
)
