# plat

A simple Go package that acts like Python's `platformdirs`, providing standard directory paths for your application on Linux, macOS, and Windows.

## Installation

```bash
go get github.com/kmuchiri/plat
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/kmuchiri/plat"
)

func main() {
	// Initialize with AppName and AppAuthor (author is mainly used on Windows)
	dirs := plat.New("MyApp", "MyCompany")

	// You can also chain WithEnsureExists to automatically create directories when accessed
	// dirs := plat.New("MyApp", "MyCompany").WithEnsureExists(true)

	// Application directories
	fmt.Println("User Data Dir:", dirs.UserDataDir())
	fmt.Println("User Config Dir:", dirs.UserConfigDir())
	fmt.Println("User Cache Dir:", dirs.UserCacheDir())
	fmt.Println("User State Dir:", dirs.UserStateDir())
	fmt.Println("User Log Dir:", dirs.UserLogDir())
	fmt.Println("User Runtime Dir:", dirs.UserRuntimeDir())
	
	// System-wide (Site) directories
	fmt.Println("Site Data Dir:", dirs.SiteDataDir())
	fmt.Println("Site Config Dir:", dirs.SiteConfigDir())
	
	// Standard User Media Directories
	fmt.Println("Documents:", dirs.UserDocumentsDir())
	fmt.Println("Downloads:", dirs.UserDownloadsDir())
	fmt.Println("Pictures:", dirs.UserPicturesDir())
}
```

## Directory Types Supported

### Application Directories (Scoped to AppName)
* `UserDataDir()`, `SiteDataDir()`
* `UserConfigDir()`, `SiteConfigDir()`
* `UserPreferenceDir()` (Distinct on macOS, otherwise same as Config)
* `UserCacheDir()`, `SiteCacheDir()`
* `UserStateDir()`, `SiteStateDir()`
* `UserLogDir()`, `SiteLogDir()`
* `UserRuntimeDir()`, `SiteRuntimeDir()`

### User Media Directories (Standard user-facing folders)
* `UserDocumentsDir()`
* `UserDownloadsDir()`
* `UserPicturesDir()`
* `UserVideosDir()`
* `UserMusicDir()`
* `UserDesktopDir()`
* `UserProjectsDir()`
* `UserPublicShareDir()`
* `UserTemplatesDir()`
* `UserFontsDir()`
* `UserBinDir()`, `SiteBinDir()`
* `UserApplicationsDir()`, `SiteApplicationsDir()`
