package plat

import (
	"os"
	"path/filepath"
)

func (p *PlatDirs) appendPaths(base string) string {
	path := base
	if p.AppAuthor != "" {
		path = filepath.Join(path, p.AppAuthor)
	}
	if p.AppName != "" {
		path = filepath.Join(path, p.AppName)
	}
	return p.processPath(path)
}

func getEnvOrDefault(envKey, defaultValue string) string {
	if val := os.Getenv(envKey); val != "" {
		return val
	}
	return defaultValue
}

func (p *PlatDirs) UserDataDir() string {
	base := getEnvOrDefault("LOCALAPPDATA", filepath.Join(getHome(), "AppData", "Local"))
	return p.appendPaths(base)
}

func (p *PlatDirs) SiteDataDir() string {
	base := getEnvOrDefault("PROGRAMDATA", "C:\\ProgramData")
	return p.appendPaths(base)
}

func (p *PlatDirs) UserConfigDir() string {
	return p.UserDataDir()
}

func (p *PlatDirs) SiteConfigDir() string {
	return p.SiteDataDir()
}

func (p *PlatDirs) UserPreferenceDir() string {
	return p.UserDataDir()
}

func (p *PlatDirs) UserCacheDir() string {
	return p.processPath(filepath.Join(p.UserDataDir(), "Cache"))
}

func (p *PlatDirs) SiteCacheDir() string {
	return p.processPath(filepath.Join(p.SiteDataDir(), "Cache"))
}

func (p *PlatDirs) UserStateDir() string {
	return p.UserDataDir()
}

func (p *PlatDirs) SiteStateDir() string {
	return p.SiteDataDir()
}

func (p *PlatDirs) UserLogDir() string {
	return p.processPath(filepath.Join(p.UserDataDir(), "Logs"))
}

func (p *PlatDirs) SiteLogDir() string {
	return p.processPath(filepath.Join(p.SiteDataDir(), "Logs"))
}

func (p *PlatDirs) UserRuntimeDir() string {
	base := getEnvOrDefault("TEMP", filepath.Join(getHome(), "AppData", "Local", "Temp"))
	return p.appendPaths(base)
}

func (p *PlatDirs) SiteRuntimeDir() string {
	return p.UserRuntimeDir()
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
	return p.processPath(getEnvOrDefault("PUBLIC", "C:\\Users\\Public"))
}

func (p *PlatDirs) UserTemplatesDir() string {
	appdata := getEnvOrDefault("APPDATA", filepath.Join(getHome(), "AppData", "Roaming"))
	return p.processPath(filepath.Join(appdata, "Microsoft", "Windows", "Templates"))
}

func (p *PlatDirs) UserFontsDir() string {
	localappdata := getEnvOrDefault("LOCALAPPDATA", filepath.Join(getHome(), "AppData", "Local"))
	return p.processPath(filepath.Join(localappdata, "Microsoft", "Windows", "Fonts"))
}

func (p *PlatDirs) UserBinDir() string {
	localappdata := getEnvOrDefault("LOCALAPPDATA", filepath.Join(getHome(), "AppData", "Local"))
	return p.processPath(filepath.Join(localappdata, "Programs"))
}

func (p *PlatDirs) SiteBinDir() string {
	programdata := getEnvOrDefault("PROGRAMDATA", "C:\\ProgramData")
	return p.processPath(filepath.Join(programdata, "bin"))
}

func (p *PlatDirs) UserApplicationsDir() string {
	appdata := getEnvOrDefault("APPDATA", filepath.Join(getHome(), "AppData", "Roaming"))
	return p.processPath(filepath.Join(appdata, "Microsoft", "Windows", "Start Menu", "Programs"))
}

func (p *PlatDirs) SiteApplicationsDir() string {
	programdata := getEnvOrDefault("PROGRAMDATA", "C:\\ProgramData")
	return p.processPath(filepath.Join(programdata, "Microsoft", "Windows", "Start Menu", "Programs"))
}
