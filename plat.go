package plat

import (
	"os"
)

// PlatDirs represents the directory configuration for an application.
type PlatDirs struct {
	AppName      string
	AppAuthor    string
	EnsureExists bool
}

// New creates a new PlatDirs instance.
// appName is the name of the application.
// appAuthor is the name of the author or organization.
func New(appName, appAuthor string) *PlatDirs {
	return &PlatDirs{
		AppName:   appName,
		AppAuthor: appAuthor,
	}
}

// WithEnsureExists configures the instance to automatically create directories when paths are requested.
func (p *PlatDirs) WithEnsureExists(ensure bool) *PlatDirs {
	p.EnsureExists = ensure
	return p
}

// processPath automatically creates the directory if EnsureExists is true.
func (p *PlatDirs) processPath(path string) string {
	if p.EnsureExists {
		_ = os.MkdirAll(path, 0755)
	}
	return path
}

func getHome() string {
	home, _ := os.UserHomeDir()
	return home
}
