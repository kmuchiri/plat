package plat

import (
	"path/filepath"
)

func (p *PlatDirs) appendPaths(base string) string {
	path := base
	if p.AppName != "" {
		path = filepath.Join(path, p.AppName)
	}
	return p.processPath(path)
}

func (p *PlatDirs) UserDataDir() string {
	return p.appendPaths(filepath.Join(getHome(), "Library", "Application Support"))
}

func (p *PlatDirs) SiteDataDir() string {
	return p.appendPaths(filepath.Join("/", "Library", "Application Support"))
}

func (p *PlatDirs) UserConfigDir() string {
	return p.UserDataDir()
}

func (p *PlatDirs) SiteConfigDir() string {
	return p.SiteDataDir()
}

func (p *PlatDirs) UserPreferenceDir() string {
	return p.appendPaths(filepath.Join(getHome(), "Library", "Preferences"))
}

func (p *PlatDirs) UserCacheDir() string {
	return p.appendPaths(filepath.Join(getHome(), "Library", "Caches"))
}

func (p *PlatDirs) SiteCacheDir() string {
	return p.appendPaths(filepath.Join("/", "Library", "Caches"))
}

func (p *PlatDirs) UserStateDir() string {
	return p.UserDataDir()
}

func (p *PlatDirs) SiteStateDir() string {
	return p.SiteDataDir()
}

func (p *PlatDirs) UserLogDir() string {
	return p.appendPaths(filepath.Join(getHome(), "Library", "Logs"))
}

func (p *PlatDirs) SiteLogDir() string {
	return p.appendPaths(filepath.Join("/", "Library", "Logs"))
}

func (p *PlatDirs) UserRuntimeDir() string {
	return p.UserDataDir()
}

func (p *PlatDirs) SiteRuntimeDir() string {
	return p.SiteDataDir()
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
	return p.processPath(filepath.Join(getHome(), "Movies"))
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
	return p.processPath(filepath.Join(getHome(), "Library", "Fonts"))
}

func (p *PlatDirs) UserBinDir() string {
	return p.processPath(filepath.Join("/", "usr", "local", "bin"))
}

func (p *PlatDirs) SiteBinDir() string {
	return p.processPath(filepath.Join("/", "usr", "local", "bin"))
}

func (p *PlatDirs) UserApplicationsDir() string {
	return p.processPath(filepath.Join(getHome(), "Applications"))
}

func (p *PlatDirs) SiteApplicationsDir() string {
	return p.processPath(filepath.Join("/", "Applications"))
}
