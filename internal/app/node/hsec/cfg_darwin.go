//go:build darwin
// +build darwin

package hsec

func reportFile() string {
	return "/opt/skynx/var/lib/report.hsr"
}

func rootTargetDir() string {
	return "/"
}

func skipDirs() []string {
	return []string{}
}

func globalCacheDir() string {
	return "/opt/skynx/var/cache"
}

/*
func globalCacheDir() string {
	tmpDir, err := os.UserCacheDir()
	if err != nil {
		tmpDir = os.TempDir()
	}

	return filepath.Join(tmpDir, "skynx", "cache")
}
*/
