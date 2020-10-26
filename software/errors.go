package software

import "fmt"

// DistroNotSupportedError will be raised when there is no custom command
// for the wanted distro
type DistroNotSupportedError struct {
	Pkg    Package
	Distro string
}

func (e *DistroNotSupportedError) Error() string {
	return fmt.Sprintf("Package %s has no custom command for distro %s", e.Pkg.Name, e.Distro)
}

// OSNotSupportedError occurrs when the script is run outside a Linux eviroment
type OSNotSupportedError struct {
	OS string
}

func (e *OSNotSupportedError) Error() string {
	return fmt.Sprintf("This Script currently does not support the OS %s", e.OS)
}
