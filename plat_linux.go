package plat

import (
	"os"
	"path/filepath"
	"strconv"
)

func getEnvOrDefault(envKey, defaultValue string) string {
	if val := os.Getenv(envKey); val != "" {
		return val
	}
	return defaultValue
}

func (p *PlatDirs) appendPaths(base string) string {
	path := base
	if p.AppName != "" {
		path = filepath.Join(path, p.AppName)
	}
	return p.processPath(path)
}

func (p *PlatDirs) UserDataDir() string {
	base := getEnvOrDefault("XDG_DATA_HOME", filepath.Join(getHome(), ".local", "share"))
	return p.appendPaths(base)
}

func (p *PlatDirs) SiteDataDir() string {
	return p.appendPaths(getEnvOrDefault("XDG_DATA_DIRS", "/usr/local/share"))
}

func (p *PlatDirs) UserConfigDir() string {
	base := getEnvOrDefault("XDG_CONFIG_HOME", filepath.Join(getHome(), ".config"))
	return p.appendPaths(base)
}

func (p *PlatDirs) SiteConfigDir() string {
	return p.appendPaths(getEnvOrDefault("XDG_CONFIG_DIRS", "/etc/xdg"))
}

func (p *PlatDirs) UserPreferenceDir() string {
	return p.UserConfigDir()
}

func (p *PlatDirs) UserCacheDir() string {
	base := getEnvOrDefault("XDG_CACHE_HOME", filepath.Join(getHome(), ".cache"))
	return p.appendPaths(base)
}

func (p *PlatDirs) SiteCacheDir() string {
	return p.appendPaths("/var/cache")
}

func (p *PlatDirs) UserStateDir() string {
	base := getEnvOrDefault("XDG_STATE_HOME", filepath.Join(getHome(), ".local", "state"))
	return p.appendPaths(base)
}

func (p *PlatDirs) SiteStateDir() string {
	return p.appendPaths("/var/lib")
}

func (p *PlatDirs) UserLogDir() string {
	return p.processPath(filepath.Join(p.UserStateDir(), "log"))
}

func (p *PlatDirs) SiteLogDir() string {
	return p.appendPaths("/var/log")
}

func (p *PlatDirs) UserRuntimeDir() string {
	uid := strconv.Itoa(os.Getuid())
	base := getEnvOrDefault("XDG_RUNTIME_DIR", filepath.Join("/run/user", uid))
	return p.appendPaths(base)
}

func (p *PlatDirs) SiteRuntimeDir() string {
	return p.appendPaths("/run")
}

// User media directories

func (p *PlatDirs) UserDocumentsDir() string {
	return p.processPath(filepath.Join(getHome(), "Documents"))
}

func (p *PlatDirs) UserDownloadsDir() string {
	return p.processPath(filepath.Join(getHome(), "Downloads"))
}

func (p *PlatDirs) UserPicturesDir() string {
	return p.processPath(filepath.Join(getHome(), "Pictures"))
}

func (p *PlatDirs) UserVideosDir() string {
	return p.processPath(filepath.Join(getHome(), "Videos"))
}

func (p *PlatDirs) UserMusicDir() string {
	return p.processPath(filepath.Join(getHome(), "Music"))
}

func (p *PlatDirs) UserDesktopDir() string {
	return p.processPath(filepath.Join(getHome(), "Desktop"))
}

func (p *PlatDirs) UserProjectsDir() string {
	return p.processPath(filepath.Join(getHome(), "Projects"))
}

func (p *PlatDirs) UserPublicShareDir() string {
	return p.processPath(filepath.Join(getHome(), "Public"))
}

func (p *PlatDirs) UserTemplatesDir() string {
	return p.processPath(filepath.Join(getHome(), "Templates"))
}

func (p *PlatDirs) UserFontsDir() string {
	return p.processPath(filepath.Join(getHome(), ".local", "share", "fonts"))
}

func (p *PlatDirs) UserBinDir() string {
	return p.processPath(filepath.Join(getHome(), ".local", "bin"))
}

func (p *PlatDirs) SiteBinDir() string {
	return p.processPath("/usr/local/bin")
}

func (p *PlatDirs) UserApplicationsDir() string {
	return p.processPath(filepath.Join(getHome(), ".local", "share", "applications"))
}

func (p *PlatDirs) SiteApplicationsDir() string {
	return p.processPath("/usr/share/applications")
}
